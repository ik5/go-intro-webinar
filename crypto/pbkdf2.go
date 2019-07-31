package crypto

import (
	"crypto/sha256"

	"golang.org/x/crypto/pbkdf2"
)

func genPBKDF2(pass string, salt []byte) ([]byte, error) {
	dk := pbkdf2.Key([]byte(pass), salt, 4096, 32, sha256.New)
	return dk, nil
}
