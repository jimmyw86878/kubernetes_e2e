package testcases

import (
	"fmt"
	"kubernetes_e2e/tests/kubeclient"
)

//Teardown is to clean test environment
func Teardown() {
	fmt.Println("====== Clean kubernetes testing pod ======")
	res, err := kubeclient.ExecKubectl("kubectl delete pod -l app=nsenter")
	fmt.Println(res)
	if err != nil {
		fmt.Println("Error: " + err.Error())
	}
}
