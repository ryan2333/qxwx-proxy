package tools

import (
	"crypto/md5"
	"encoding/hex"
)

func GeneratorMd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
