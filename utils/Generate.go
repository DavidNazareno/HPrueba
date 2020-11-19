package utils

import (
	"crypto/rand"
	"fmt"
)

//Generate generador texto aleatorio para los token y tickets
func Generate(size int) string {
	b := make([]byte, size)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
