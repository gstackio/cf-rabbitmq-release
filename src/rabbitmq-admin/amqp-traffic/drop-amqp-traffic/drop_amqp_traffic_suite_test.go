package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDropAmqpTraffic(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DropAmqpTraffic Suite")
}
