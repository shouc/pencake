package utils

import "os"

func IsFileExist(path string) bool {
	info, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
