package main

import (
	"flag"
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

// Used for decryption
type DATA_BLOB struct {
	cbData uint32
	pbData *byte
}

func NewBlob(d []byte) *DATA_BLOB {
	if len(d) == 0 {
		return &DATA_BLOB{}
	}
	return &DATA_BLOB{
		pbData: &d[0],
		cbData: uint32(len(d)),
	}
}

func (b *DATA_BLOB) ToByteArray() []byte {
	d := make([]byte, b.cbData)
	copy(d, (*[1 << 30]byte)(unsafe.Pointer(b.pbData))[:])
	return d
}

func usage() {
	fmt.Println("Usage: ./d [encrypted_text]")
	flag.PrintDefaults()
	os.Exit(2)
}

func decrypt(data string) string {
	dllcrypt32 := syscall.NewLazyDLL("Crypt32.dll")
	dllkernel32 := syscall.NewLazyDLL("Kernel32.dll")

	procDecryptData := dllcrypt32.NewProc("CryptUnprotectData")
	procLocalFree := dllkernel32.NewProc("LocalFree")
	var outblob DATA_BLOB
	r, _, _ := procDecryptData.Call(uintptr(unsafe.Pointer(NewBlob([]byte(data)))), 0, 0, 0, 0, 0,
		uintptr(unsafe.Pointer(&outblob)))
	if r == 0 {
		return ""
	}
	defer procLocalFree.Call(uintptr(unsafe.Pointer(outblob.pbData)))
	return string(outblob.ToByteArray())
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Encrypted text missing")
		os.Exit(1)
	}
	fmt.Printf(decrypt(args[0]))
}
