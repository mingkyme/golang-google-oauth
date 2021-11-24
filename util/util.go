package util

import (
	"crypto/rand"
	"encoding/base64"
)

func RandomString(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	return state
}
