package encryption

import (
	"testing"
)

func TestEncryption(t *testing.T) {
	password := "password"
	hash, err := GenerateHash(password)
	if err != nil {
		t.Errorf("GenerateHash() Failed error %s ", err.Error())
	}
	err = CheckPasswordHash(password, hash)
	if err != nil {
		t.Errorf("CheckPasswordHash() Failed error %s ", err.Error())
	}
}
