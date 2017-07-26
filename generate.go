package main

import (
	"crypto/sha512"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

func Generate(master []byte, pi *PasswordInfo) (string, error) {
	if pi.Salt == nil {
		return "", errors.New("No salt! Cannot generate password.")
	}
	combined := append(master, []byte(pi.Name)...)
	dk := pbkdf2.Key(combined, pi.Salt, 2000, 32, sha512.New)

	str := EncodeLettersNumbersSpecial(dk)
	fmt.Println("full string", str)
	return shortenPw(str, pi)
}

func shortenPw(pw string, pi *PasswordInfo) (string, error) {
	var start int
	for {
		end := start + pi.Length
		if end > len(pw)-1 {
			return "", errors.New("Cannot satisfy requirements")
		}
		shortened := pw[start:end]
		if validPw(shortened, pi.RequiredChars) {
			return shortened, nil
		}
		start++
	}
}

func validPw(pw string, rcs []*RequiredChars) bool {
	for _, rc := range rcs {
		if !validRequiredChar(pw, rc) {
			return false
		}
	}
	return true
}

func validRequiredChar(pw string, rc *RequiredChars) bool {
	var found int
	for i := 0; i < len(rc.Chars); i++ {
		char := string(rc.Chars[i])
		if strings.Contains(pw, char) {
			found++
		}
	}
	return found >= rc.Num
}
