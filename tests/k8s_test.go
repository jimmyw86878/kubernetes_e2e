package k8s_test

import (
	"fmt"
	tc "kubernetes_e2e/tests/testcases"
	"kubernetes_e2e/tests/util"

	"github.com/onsi/ginkgo"
)

var _ = ginkgo.Describe("K8S: ", func() {
	// Test case 1
	ginkgo.Context("Testing k8s health", func() {
		ginkgo.It("should return normal output for all nodes", tc.HealthCheckTest)
	})
	//Test case 2
	ginkgo.Context("Testing test_pod_fail yaml", func() {
		ginkgo.It("should return failure", tc.FailPodTest)
	})
	// Test case 3
	ginkgo.Context("Testing test_pod_success yaml", func() {
		ginkgo.It("should return success", tc.SuccessPodTest)
	})
	// Test case 4
	ginkgo.Context("Testing reboot a node", func() {
		ginkgo.It("should success to join back to cluster", tc.RebootNodeTest)//default is to restart minikube node
	})
})

var _ = ginkgo.AfterSuite(func() {
	tc.Teardown()
})