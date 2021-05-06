package kubeclient

import (
	"fmt"
	"kubernetes_e2e/tests/util"
)

//ExecKubectl define
func ExecKubectl(cmd string) (string, error) {
	kubecmd := fmt.Sprintf("docker exec kubeclient %s", cmd)
	res, err := util.Exec(kubecmd)
	if err != nil {
		return "", err
	}
	return res, nil
}

//ExecTestYaml define
func ExecTestYaml(filename string) bool {
	kubecmd := fmt.Sprintf("docker exec kubeclient seal autonomous --policy-file /home/testyaml/%s", filename)
	return util.ExecYaml(kubecmd)
}
