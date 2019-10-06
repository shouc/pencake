package utils

func WindowsDecrypt(data string) string {
	result := RunCommand("windows\\d.exe " + data)
	return result
}
