package util

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
)

var (
	defaultTimeout = 30 * time.Second
)

//Exec to execute command in container, default timeout is 30 seconds
func Exec(cmd string, timeout ...time.Duration) (string, error) {
	execTimeout := defaultTimeout
	if len(timeout) != 0 {
		execTimeout = timeout[0]
	}
	ctx, cancel := context.WithTimeout(context.Background(), execTimeout)
	defer cancel()
	res, err := exec.CommandContext(ctx, "/bin/bash", "-c", cmd).Output()
	if ctx.Err() == context.DeadlineExceeded {
		return "Command timed out", fmt.Errorf("Command time out")
	}
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
