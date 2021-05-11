package testcases

import (
	"fmt"
	"kubernetes_e2e/tests/kubeclient"
	ct "kubernetes_e2e/tests/constants"
	"kubernetes_e2e/tests/util"
	"time"

	"github.com/onsi/gomega"
)

//RebootNodeTest is to reboot a node
func RebootNodeTest() {
	fmt.Println("====== Test reboot ======")
	fmt.Printf("Rebooting %s\n", ct.RebootNode)
	err := kubeclient.ExecReboot(ct.RebootNode)
	gomega.Expect(err).To(gomega.BeNil())
	//since minikube won't restart after reboot, start it manually
	//assume minikube should start in 3 minutes
	res, err := util.Exec("minikube start", 180*time.Second)
	fmt.Println(res)
	gomega.Expect(err).To(gomega.BeNil())
	fmt.Println("====== Test reboot END ======")
}
