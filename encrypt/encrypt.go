package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	"mypass/data"

	"golang.org/x/crypto/pbkdf2"
)

var saltSize = 12
var keySize = 32

// Encrypt takes the plain-text password and the encryption key, and
// returns the encrypted password and the salt used to encrypt it
func Encrypt(pw data.Secret, salt []byte, key data.Secret) (encPw []byte,
	err error) {

	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCMWithNonceSize(c, len(salt))
	if err != nil {
		return nil, err
	}

	return gcm.Seal(nil, salt, pw, nil), nil
}

func Decrypt(encPw, salt []byte, key data.Secret) (pw []byte, err error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCMWithNonceSize(c, len(salt))
	if err != nil {
		return nil, err
	}

	return gcm.Open(nil, salt, encPw, nil)
}

func MasterToKey(master data.Secret, salt []byte) []byte {
	return pbkdf2.Key(append(master, salt...), salt, 4096, keySize, sha512.New)
}

func NewSalt(len int) ([]byte, error) {
	b := make([]byte, len)

	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
