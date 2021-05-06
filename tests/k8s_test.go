package k8s_test

import (
	"fmt"
	tc "kubernetes_e2e/tests/testcases"
	"kubernetes_e2e/tests/util"

	"github.com/onsi/ginkgo"
)

var (
	testcase1 = "Testing k8s health"
	testcase2 = "Testing test_pod_fail yaml"
	testcase3 = "Testing test_pod_success yaml"
)

var _ = ginkgo.Describe("K8S: ", func() {
	// Test case 1
	ginkgo.Context(testcase1, func() {
		ginkgo.It("should return normal output for all nodes", func() {
			if !util.CheckRunOrNot(testcase1) {
				ginkgo.Skip(fmt.Sprintf("Skip %s", testcase1))
			}
			tc.HealthCheckTest()
		})
	})
	//Test case 2
	ginkgo.Context(testcase2, func() {
		ginkgo.It("should return failure", func() {
			if !util.CheckRunOrNot(testcase2) {
				ginkgo.Skip(fmt.Sprintf("Skip %s", testcase2))
			}
			tc.FailPodTest()
		})
	})
	// Test case 3
	ginkgo.Context(testcase3, func() {
		ginkgo.It("should return success", func() {
			if !util.CheckRunOrNot(testcase3) {
				ginkgo.Skip(fmt.Sprintf("Skip %s", testcase3))
			}
			tc.SuccessPodTest()
		})
	})
})
