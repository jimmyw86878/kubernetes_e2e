package kubeclient

import (
	"fmt"
	"kubernetes_e2e/tests/util"
)

var (
	prefixKubeClient = "docker exec kubeclient"
)

//ExecKubectl is to execute kubectl command to target kubernetes cluster
func ExecKubectl(cmd string) (string, error) {
	kubecmd := fmt.Sprintf("%s %s", prefixKubeClient, cmd)
	res, err := util.Exec(kubecmd)
	if err != nil {
		return "", err
	}
	return res, nil
}

//ExecTestYaml is to execute Powerfulseal yamls
func ExecTestYaml(filename string) bool {
	kubecmd := fmt.Sprintf("%s seal autonomous --policy-file /home/testyaml/%s", prefixKubeClient, filename)
	return util.ExecYaml(kubecmd)
}

//ExecReboot is to reboot a node in kubernetes cluster
func ExecReboot(node string) error {
	kubecmd := fmt.Sprintf("%s kubectl node-shell %s -- sudo reboot &> /dev/null", prefixKubeClient, node)
	res, err := util.Exec(kubecmd)
	//reboot may cause timeout, return nil error if reboot command timed out
	if err != nil && res != "Command timed out" {
		return err
	}
	return nil
}
