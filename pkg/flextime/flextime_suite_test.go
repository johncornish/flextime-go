package flextime_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFlextime(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Flextime Suite")
}
