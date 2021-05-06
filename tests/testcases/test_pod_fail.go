package testcases

import (
	"fmt"
	"kubernetes_e2e/tests/kubeclient"

	"github.com/onsi/gomega"
)

//FailPodTest is an example for deleting a pod in kube-system and checking should be failed
func FailPodTest() {
	fmt.Println("====== FailPodTest ======")
	//test reschedule_kubemonitor.yml
	result := kubeclient.ExecTestYaml("test_pod_fail.yml")
	gomega.Expect(result).NotTo(gomega.BeTrue())
	fmt.Println("====== FailPodTest END ======")
}
