package testcases

import (
	"fmt"
	"kubernetes_e2e/tests/kubeclient"

	"github.com/onsi/gomega"
)

//HealthCheckTest is to check health of k8s cluster
func HealthCheckTest() {
	fmt.Println("====== Test health ======")
	res, err := kubeclient.ExecKubectl("kubectl get node")
	fmt.Println(res)
	gomega.Expect(res).NotTo(gomega.Equal(""))
	gomega.Expect(err).To(gomega.BeNil())
	fmt.Println("====== Test health END ======")
}
