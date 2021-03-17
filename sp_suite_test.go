package sp_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sp Suite")
}
