package crypto

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// GenSalt creates a random crypto based salt, based on a given size.
// if not given, then DefaultSaltSize is in use
func GenSalt(size int) []byte {
	s := size
	if s == 0 {
		s = DefaultSaltSize
	}

	salt := make([]byte, s)
	rand.Read(salt)
	return salt
}

// GenPassword generate a new password string
func GenPassword(cryptoType int, str string, salt []byte) (string, error) {
	var dk []byte
	var err error
	switch cryptoType {
	case SCrypt:
		dk, err = genSCrypt(str, salt)
	case BCrypt:
		// TODO: implement it from golang.org/x/crypto/bcrypt
	case Argon2:
		dk, err = genArgon2(str, salt)
	case PBKDF2:
		dk, err = genPBKDF2(str, salt)
	default:
		err = fmt.Errorf("cryptoType %d is not supported", cryptoType)
	}
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%04x$%x$%x", cryptoType, salt, dk), nil
}

// IsValidPassword Validates a given password with the result of the string
func IsValidPassword(pass, str string) (bool, error) {
	elements := strings.Split(str, "$")
	if len(elements) != 3 {
		return false, errors.New("Invalid encrypted password ")
	}
	cryptoType, err := strconv.Atoi(elements[0])
	if err != nil {
		return false, err
	}
	if cryptoType == 0 {
		return false, errors.New("Invalid encrypted password")
	}
	salt, err := hex.DecodeString(elements[1])
	if err != nil {
		return false, err
	}
	password, err := hex.DecodeString(elements[2])
	if err != nil {
		return false, err
	}

	newPass, err := GenPassword(cryptoType, pass, salt)
	if err != nil {
		return false, nil
	}
	elements = strings.Split(newPass, "$")
	newPassword, err := hex.DecodeString(elements[2])
	return bytes.Equal(password, newPassword), nil
}
