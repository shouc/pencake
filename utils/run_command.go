package utils

import (
	"fmt"
	"os/exec"
)

func RunCommand(command string) string {
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("[+] Command failed with %s\n", err)
		return ""
	}
	// ripe away \n
	return string(output[:len(output)-1])
}
