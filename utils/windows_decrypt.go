package utils

func WindowsDecrypt(data string) string {
	result := RunCommand("./windows/d " + data)
	return result
}
