package main

import (
	"crypto/rand"
	"errors"
)

func GenerateSalt(master []byte, pi *PasswordInfo) error {
	for i := 0; i < 5; i++ {
		salt, err := newSalt(16)
		if err != nil {
			return err
		}
		pi.Salt = salt
		if _, err := Generate(master, pi); err == nil {
			return nil
		}
	}
	pi.Salt = nil
	return errors.New("Cannot generate salt that meets requirements")
}

func newSalt(len int) ([]byte, error) {
	b := make([]byte, len)

	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
