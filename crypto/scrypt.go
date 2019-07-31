package crypto

import "golang.org/x/crypto/scrypt"

// generate a SCrypt encryption for password
func genSCrypt(str string, salt []byte) ([]byte, error) {
	dk, err := scrypt.Key([]byte(str), salt, 1<<15, 8, 1, 32)
	return dk, err
}
