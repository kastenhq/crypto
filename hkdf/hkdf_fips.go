//go:build goexperiment.opensslcrypto

package hkdf

import (
	"hash"
	"io"

	_ "unsafe"
)

//go:linkname hkdfExtract vendor/github.com/golang-fips/openssl/v2.ExtractHKDF
func hkdfExtract(h func() hash.Hash, secret, salt []byte) ([]byte, error)

//go:linkname hkdfExpand vendor/github.com/golang-fips/openssl/v2.ExpandHKDF
func hkdfExpand(h func() hash.Hash, pseudorandomKey, info []byte) (io.Reader, error)
