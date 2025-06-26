//go:build go1.24

package hkdf

import (
	"bytes"
	"crypto/hkdf"
	"hash"
	"io"
)

func hkdfExtract(h func() hash.Hash, secret, salt []byte) ([]byte, error) {
	return hkdf.Extract(h, secret, salt)
}

func hkdfExpand(h func() hash.Hash, pseudorandomKey, info []byte) (io.Reader, error) {
	// Since the golang.org/x/crypto signature doesn't use a key length
	// we use the size of the hash
	hashSize := h().Size()
	b, err := hkdf.Expand(h, string(pseudorandomKey), info, hashSize)
	return bytes.NewReader(b), err
}
