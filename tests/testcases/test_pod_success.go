package testcases

import (
	"fmt"
	"kubernetes_e2e/tests/kubeclient"

	"github.com/onsi/gomega"
)

//SuccessPodTest is an example for deleting a pod in kube-system and checking should success
func SuccessPodTest() {
	fmt.Println("====== FailPodTest ======")
	//test reschedule_kubemonitor.yml
	result := kubeclient.ExecTestYaml("test_pod_success.yml")
	gomega.Expect(result).To(gomega.BeTrue())
	fmt.Println("====== FailPodTest END ======")
}
