package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(source string) string {
	h := md5.New()
	h.Write([]byte(source))
	return hex.EncodeToString(h.Sum(nil))
}
