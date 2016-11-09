package federator

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

// AMQPClient is a struct containing an AMQP connection
type AMQPClient struct {
	conn *amqp.Connection
	done chan error
}

// Initialize makes the initial connection to the AMQP server
func (a *AMQPClient) Initialize(address string) (err error) {
	for i := 1; i <= 5; i++ {
		a.conn, err = amqp.Dial(address)
		if err == nil {
			return nil
		}
		log.Printf("Could not connect due to: %s, retrying in 5 seconds...\n", err)
		time.Sleep(5 * time.Second)
	}

	a.done = make(chan error)
	return err
}

// Close closes the  AMQP server
func (a *AMQPClient) Close() (err error) {
	if a.conn != nil {
		return a.conn.Close()
	}
	return nil
}

// AmqpSend sends a message to the specified exchange
func (a *AMQPClient) AmqpSend(exchange, exchangeType, routingKey, body, correlationID, replyTo string) error {
	log.Printf("got Connection, getting Channel")
	channel, err := a.conn.Channel()
	if err != nil {
		return fmt.Errorf("Channel: %s", err)
	}

	log.Printf("got Channel, declaring %q Exchange (%q)", exchangeType, exchange)
	err = channel.ExchangeDeclare(
		exchange,     // name
		exchangeType, // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // noWait
		nil,          // arguments
	)
	if err != nil {
		return fmt.Errorf("Exchange Declare: %s", err)
	}

	log.Printf("declared Exchange, publishing %dB body (%q)", len(body), body)
	if err := channel.Publish(
		exchange,   // publish to an exchange
		routingKey, // routing to 0 or more queues
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			CorrelationId:   correlationID,
			ReplyTo:         replyTo,
			Body:            []byte(body),
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,              // 0-9
		},
	); err != nil {
		return fmt.Errorf("Exchange Publish: %s", err)
	}

	return nil
}

// AmqpListen connects to an AMQP exchange and returns a channel that will contain messages
func (a *AMQPClient) AmqpListen(exchange, exchangeType, queueName, bindingKey, consumerTag string) (<-chan amqp.Delivery, error) {
	var err error

	log.Printf("got Connection, getting Channel")
	channel, err := a.conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("Channel: %s", err)
	}

	log.Printf("got Channel, declaring Exchange (%q)", exchange)
	if err = channel.ExchangeDeclare(
		exchange,     // name of the exchange
		exchangeType, // type
		true,         // durable
		false,        // delete when complete
		false,        // internal
		false,        // noWait
		nil,          // arguments
	); err != nil {
		return nil, fmt.Errorf("Exchange Declare: %s", err)
	}

	log.Printf("declared Exchange, declaring Queue %q", queueName)
	queue, err := channel.QueueDeclare(
		queueName, // name of the queue
		true,      // durable
		false,     // delete when usused
		true,      // exclusive
		false,     // noWait
		nil,       // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("Queue Declare: %s", err)
	}

	log.Printf("declared Queue (%q %d messages, %d consumers), binding to Exchange (key %q)",
		queue.Name, queue.Messages, queue.Consumers, bindingKey)

	if err = channel.QueueBind(
		queue.Name, // name of the queue
		bindingKey, // bindingKey
		exchange,   // sourceExchange
		false,      // noWait
		nil,        // arguments
	); err != nil {
		return nil, fmt.Errorf("Queue Bind: %s", err)
	}

	log.Printf("Queue bound to Exchange, starting Consume (consumer tag %q)", consumerTag)
	deliveries, err := channel.Consume(
		queue.Name,  // name
		consumerTag, // consumerTag,
		false,       // noAck
		false,       // exclusive
		false,       // noLocal
		false,       // noWait
		nil,         // arguments
	)
	if err != nil {
		return nil, fmt.Errorf("Queue Consume: %s", err)
	}

	log.Printf("**Now listening on exchange: %s**\n", exchange)

	return deliveries, nil
}
