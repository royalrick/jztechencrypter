package encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(raw []byte) string {
	h := md5.New()
	h.Write(raw)
	return hex.EncodeToString(h.Sum(nil))
}
