package channeliterator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestChannelIterator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ChannelIterator Suite")
}
