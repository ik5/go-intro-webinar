package crypto

import "golang.org/x/crypto/argon2"

func genArgon2(pass string, salt []byte) ([]byte, error) {
	key := argon2.IDKey([]byte(pass), salt, 1, 64*1024, 4, 32)
	return key, nil
}
