package utils

import (
	"encoding/base64"
)

func Base64Encode(content string) string {
	data := []byte(content)
	return base64.StdEncoding.EncodeToString(data)
}
