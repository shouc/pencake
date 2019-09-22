package utils

import (
	"fmt"
	"testing"
)

func TestRunCommand(t *testing.T) {
	output := RunCommand("echo 'hi'")
	fmt.Println(output == "hi")
	if output != "hi" {
		t.Fail()
	}
}
