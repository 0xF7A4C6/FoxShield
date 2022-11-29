package utils

import (
	"fmt"
	"math/rand"
)

func RandomStr(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}
