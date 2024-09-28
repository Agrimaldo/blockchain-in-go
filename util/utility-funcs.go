package util

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func NewHash() string {
	data := make([]byte, 75)
	_, err := io.ReadFull(rand.Reader, data)
	if err != nil {
		panic(fmt.Sprint("Error reading random bytes:", err))
	}

	return base64.StdEncoding.EncodeToString(data)
}
