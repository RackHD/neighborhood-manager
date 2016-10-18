package federator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestAmqp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Amqp Suite")
}
