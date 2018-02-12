package test

import "os/exec"

func HasDockerInstalled() bool {
	return HasInstalled("docker")
}

func HasInstalled(program string) bool {
	err := exec.Command(program).Run()
	if err != nil {
		return false
	} else {
		return true
	}
}
