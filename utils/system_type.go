package utils

import "runtime"

func SystemType() string {
	return runtime.GOOS
}
