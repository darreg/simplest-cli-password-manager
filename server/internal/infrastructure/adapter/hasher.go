package adapter

import (
	"crypto/sha256"
	"fmt"
)

type Hasher struct{}

func NewHasher() *Hasher {
	return &Hasher{}
}

func (h *Hasher) Hash(password string) string {
	pwd := sha256.New()
	pwd.Write([]byte(password))

	return fmt.Sprintf("%x", pwd.Sum(nil))
}
