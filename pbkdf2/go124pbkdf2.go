//go:build go1.24

package pbkdf2

import (
	"crypto/pbkdf2"
	"hash"
)

func pbkdf2_key_derivation(password, salt []byte, iter, keyLen int, h func() hash.Hash) ([]byte, error) {
	return pbkdf2.Key(h, string(password), salt, iter, keyLen)
}
