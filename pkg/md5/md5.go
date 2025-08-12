package md5

import (
	"crypto/md5"
	"encoding/hex"
)

// 将数据进行MD5加密成32位16进制
func Sum(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil)) //32位16进制
}
