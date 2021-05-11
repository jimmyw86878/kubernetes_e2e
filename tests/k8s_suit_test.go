package k8s_test

import (
	"flag"
	"testing"

	ct "kubernetes_e2e/tests/constants"

	"github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	"github.com/onsi/gomega"
)

func init() {
	flag.StringVar(&ct.RebootNode, "rebootNode", "minikube", "")
}

func TestK8SE2E(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	junitReporter := reporters.NewJUnitReporter("../junit.xml")
	ginkgo.RunSpecsWithDefaultAndCustomReporters(t, "K8S Suite", []ginkgo.Reporter{junitReporter})
}
