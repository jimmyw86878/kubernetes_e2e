package testcases

import (
	"fmt"
	"kubernetes_e2e/tests/kubeclient"
	"kubernetes_e2e/tests/util"
	"time"

	"github.com/onsi/gomega"
)

//RebootMinikubeTest is to reboot minikube
func RebootMinikubeTest() {
	fmt.Println("====== Test reboot ======")
	err := kubeclient.ExecReboot("minikube")
	gomega.Expect(err).To(gomega.BeNil())
	//since minikube won't restart after reboot, start it manually
	//assume minikube should start in 3 minutes
	res, err := util.Exec("minikube start", 180*time.Second)
	fmt.Println(res)
	gomega.Expect(err).To(gomega.BeNil())
	fmt.Println("====== Test reboot END ======")
}
