package federator_test

import (
	"github.com/RackHD/neighborhood-manager/rackhd/amqp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Federator", func() {
	var (
		amqpURI  string
		mgmtPort string
		ctag     string
		list     []federator.ExchangeConfig
	)
	BeforeEach(func() {
		amqpURI = "amqp://guest:guest@localhost:5672"
		mgmtPort = "8080"
		ctag = "nm"
		e1 := federator.ExchangeConfig{
			ExchangeName: "on.events",
			ExchangeType: "topic",
			Durable:      true,
		}
		e2 := federator.ExchangeConfig{
			ExchangeName: "on.heartbeat",
			ExchangeType: "topic",
			Durable:      true,
		}
		list = append(list, e1, e2)
	})

	Describe("NewFederator", func() {
		It("should create a new one and return a new struct", func() {
			a, err := federator.NewAmqpFed(amqpURI, mgmtPort, ctag, list)
			Expect(a).To(BeAssignableToTypeOf(&federator.AmqpFed{}))
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Describe("AddRackHD", func() {
		It("should add a RackHD node to federation", func() {
			a, err := federator.NewAmqpFed(amqpURI, mgmtPort, ctag, list)
			a.AddRackHD("10.240.18.201:9090")
			Expect(a).To(BeAssignableToTypeOf(&federator.AmqpFed{}))
			Expect(err).ToNot(HaveOccurred())
		})
	})
})
