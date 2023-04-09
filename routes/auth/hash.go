package auth

import (
	"crypto/sha256"
	"encoding/hex"
)

func Hash(A string) string {
	h := sha256.New()
	h.Write([]byte(A))
	bs := h.Sum([]byte(`epicadidash`))
	hx := hex.EncodeToString([]byte(bs))
	return string(hx)
}
