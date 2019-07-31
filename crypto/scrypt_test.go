package crypto

import (
	"bytes"
	"testing"
)

const (
	text               = "test"
	text2              = "test2"
	validEncryptionStr = "28b229f5b92c0d8379083f816985a627675706ba1cf9d5cd785f9f32e77e56f2"
)

var (
	salt            = []byte("1")
	validEncryption = []byte{
		40, 178, 41, 245, 185, 44, 13, 131, 121, 8, 63, 129, 105, 133, 166, 39, 103,
		87, 6, 186, 28, 249, 213, 205, 120, 95, 159, 50, 231, 126, 86, 242,
	}
	invalidEncryption = []byte{
		45, 178, 41, 245, 185, 44, 13, 131, 121, 8, 63, 129, 105, 133, 166, 39, 103,
		87, 6, 186, 28, 249, 213, 205, 120, 95, 159, 50, 231, 126, 86, 242,
	}
)

func TestGenSCryptValid(t *testing.T) {
	encrypted, err := genSCrypt(text, salt)
	if err != nil {
		t.Errorf("Expected encrypted data, but got error: %s", err)
	}

	compared := bytes.Compare(encrypted, validEncryption)
	if compared != 0 {
		t.Errorf("Expected %+v == %+v, got %d instead", encrypted, validEncryption, compared)
	}
}

func TestGenSCryptInValid(t *testing.T) {
	encrypted, err := genSCrypt(text, salt)
	if err != nil {
		t.Errorf("Expected encrypted data, but got error: %s", err)
	}

	compared := bytes.Compare(encrypted, invalidEncryption)
	if compared == 0 {
		t.Errorf("Expected %+v != %+v, got %d instead", encrypted, validEncryption, compared)
	}
}
