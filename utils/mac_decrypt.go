package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
)

const iv = "20202020202020202020202020202020"

func MacDecrypt(key string, enc string) string {
	if len(enc) < 3 {
		return ""
	}
	enc = Base64Encode(enc[3:])
	src := pbkdf2.Key([]byte(key), []byte("saltysalt"), 1003, 32, sha1.New)[0:16]
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	output := RunCommand(
		fmt.Sprintf("openssl enc -base64 -d -aes-128-cbc -iv '%s' -K %s <<< %s 2>/dev/null", iv, dst, enc))
	return output
}
