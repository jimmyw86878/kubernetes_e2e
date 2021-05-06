package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

//Exec to execute command in container
func Exec(cmd string) (string, error) {
	res, err := exec.Command("/bin/bash", "-c", cmd).Output()
	if err != nil {
		return "", fmt.Errorf("Error when executing command %s, err: %s", cmd, err.Error())
	}
	return string(res), nil
}

//ExecYaml to execute command in container
func ExecYaml(cmd string) bool {
	errArr := make([]error, 0)
	c := exec.Command("/bin/bash", "-c", cmd)
	stderr, err := c.StderrPipe()
	if err != nil {
		errArr = append(errArr, err)
	}
	if err := c.Start(); err != nil {
		errArr = append(errArr, err)
	}
	slurp, _ := ioutil.ReadAll(stderr)
	fmt.Printf("%s\n", slurp)
	if err := c.Wait(); err != nil {
		errArr = append(errArr, err)
	}
	if len(errArr) == 0 {
		return true
	}
	return false
}

//CheckRunOrNot is to check this testcase should run or not. Based on environment variable `GINKGO_TESTCASE`
func CheckRunOrNot(testname string) bool {
	if val, exists := os.LookupEnv("GINKGO_TESTCASE"); exists {
		if val == testname {
			return true
		}
		return false
	}
	return true
}
