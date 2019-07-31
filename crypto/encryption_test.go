package crypto

import (
	"math"
	"strings"
	"testing"
)

const (
	validGenPassword = "0001$31$28b229f5b92c0d8379083f816985a627675706ba1cf9d5cd785f9f32e77e56f2"
)

func TestGenPasswordWrongCryptoType(t *testing.T) {
	_, err := GenPassword(math.MaxInt16, text, salt)
	if err == nil {
		t.Error("Unexpected nil err")
	}
}

func TestGenPasswordValidCryptoType(t *testing.T) {
	if testing.Short() {
		t.Skip("Computation of password can take few seconds, skipping on short tests")
	}
	_, err := GenPassword(SCrypt, text, salt)
	if err != nil {
		t.Errorf("Unexpected err was provided: %s", err)
	}
}

func TestValidGenPasswordStr(t *testing.T) {
	if testing.Short() {
		t.Skip("Computation of password can take few seconds, skipping on short tests")
	}
	password, err := GenPassword(SCrypt, text, salt)
	if err != nil {
		t.Errorf("Unexpected err was provided: %s", err)
		return
	}

	compared := strings.Compare(password, validGenPassword)
	if compared != 0 {
		t.Errorf("Expected %s == %s (%d)", password, validGenPassword, compared)
	}
}

func TestInvalidGenPasswordStr(t *testing.T) {
	if testing.Short() {
		t.Skip("Computation of password can take few seconds, skipping on short tests")
	}
	password, err := GenPassword(SCrypt, text2, salt)
	if err != nil {
		t.Errorf("Unexpected err was provided: %s", err)
		return
	}

	compared := strings.Compare(password, validGenPassword)
	if compared == 0 {
		t.Errorf("Expected %s != %s (%d)", password, validGenPassword, compared)
	}
}

func TestIsValidPasswordValidError(t *testing.T) {
	_, err := IsValidPassword(text, validEncryptionStr)
	if err == nil {
		t.Error("Expected, err not to be nil")
	}
}

func TestIsValidPasswordValidErrorAndFalse(t *testing.T) {
	result, err := IsValidPassword(text, validEncryptionStr)
	if err == nil {
		t.Error("Expected, err not to be nil")
		return
	}

	if result {
		t.Error("Expected result to be false")
	}
}

func TestIsValidPasswordValidTrue(t *testing.T) {
	if testing.Short() {
		t.Skip("Computation of password can take few seconds, skipping on short tests")
	}
	result, err := IsValidPassword(text, validGenPassword)
	if err != nil {
		t.Errorf("Unexpected err was provided: %s", err)
		return
	}

	if !result {
		t.Error("Expected valid passwords, returned false")
	}
}

func TestIsValidPasswordValidFalse(t *testing.T) {
	if testing.Short() {
		t.Skip("Computation of password can take few seconds, skipping on short tests")
	}
	result, err := IsValidPassword(text2, validGenPassword)
	if err != nil {
		t.Errorf("Unexpected err was provided: %s", err)
		return
	}

	if result {
		t.Error("Expected valid passwords, returned true")
	}
}

func TestGenSaltValidLen(t *testing.T) {
	b := GenSalt(1)

	if len(b) != 1 {
		t.Errorf("Expected salt length of 1 - got %d", len(b))
	}
}

func TestGenSaltDefaultSize(t *testing.T) {
	b := GenSalt(0)

	if len(b) != DefaultSaltSize {
		t.Errorf("Expected len of %d, got %d", DefaultSaltSize, len(b))
	}
}

func TestGenSaltNegativeSize(t *testing.T) {
	b := GenSalt(-1)

	if len(b) != DefaultSaltSize {
		t.Errorf("Expected len of %d, got %d", DefaultSaltSize, len(b))
	}
}
