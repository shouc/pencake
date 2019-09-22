package utils

import (
	"io"
	"os"
)

func CopyFile(src string, dst string) {
	source, _ := os.Open(src)
	defer source.Close()
	destination, _ := os.Create(dst)
	defer destination.Close()
	_, _ = io.Copy(destination, source)
}
