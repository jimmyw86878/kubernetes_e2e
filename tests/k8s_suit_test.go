package k8s_test

import (
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	"github.com/onsi/gomega"
)

func TestK8SE2E(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	junitReporter := reporters.NewJUnitReporter("../junit.xml")
	ginkgo.RunSpecsWithDefaultAndCustomReporters(t, "K8S Suite", []ginkgo.Reporter{junitReporter})
}
