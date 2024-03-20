//go:build fipsonly

package pbkdf2

import (
	"hash"

	opensslpbkdf "github.com/golang-fips/openssl/v2"
	_ "unsafe"
)

//go:linkname pbkdf2_key_derivation vendor/github.com/golang-fips/openssl/v2.PBKDF2
func pbkdf2_key_derivation(password, salt []byte, iter, keyLen int, h func() hash.Hash) ([]byte, error) {
	return opensslpbkdf.PBKDF2(password, salt, iter, keyLen, h)
}
