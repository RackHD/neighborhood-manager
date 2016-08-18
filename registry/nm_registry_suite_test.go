package registry_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestNmRegistry(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "NmRegistry Suite")
}
