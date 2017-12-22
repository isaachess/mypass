package generate

import (
	"bytes"
	"crypto/rand"
	"errors"
	"math/big"
	"mypass/data"
)

var (
	lettersUpper          = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lettersLower          = []byte("abcdefghijklmnopqrstuvwxyz")
	numbers               = []byte("1234567890")
	symbols               = []byte("!@#$%&*()_+-")
	allLetters            = append(lettersUpper, lettersLower...)
	lettersNumbers        = append(allLetters, numbers...)
	lettersNumbersSymbols = append(lettersNumbers, symbols...)
)

// Chars represents characters to use in a password. The chars attribute is a
// set of characters to include in the random password selection. minRequired
// specifies the minimum number of this charset that must be in the final
// password. If minRequired == 0, it will not guarantee that any of these chars
// are used in the final password.
type Chars struct {
	chars       []byte
	minRequired int
}

// Common types of passwords:
// - Pure random, no requirements
// - Require upper-case
// - Require upper, lower, number, and symbol

var (
	AllChars    = []Chars{NewChars(lettersNumbersSymbols, 0)}
	StrictChars = []Chars{
		NewChars(lettersUpper, 1),
		NewChars(lettersLower, 1),
		NewChars(numbers, 1),
		NewChars(symbols, 1),
	}
)

func NewChars(chars []byte, minRequired int) Chars {
	return Chars{chars: chars, minRequired: minRequired}
}

func Generate(length int, chars []Chars) (data.Secret, error) {
	// attempt 10 times to generate a password with the required chars
	for i := 0; i < 10; i++ {
		final, err := generateBytes(length, charsetFromChars(chars))
		if err != nil {
			return nil, err
		}
		if passesChars(final, chars) {
			return final, nil
		}
	}
	return nil, errors.New("Cannot generate bytes that meet chars requirements")
}

func generateBytes(length int, charset []byte) (data.Secret, error) {
	final := make(data.Secret, length)
	charLen := big.NewInt(int64(len(charset)))
	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, charLen)
		if err != nil {
			return nil, err
		}
		final[i] = charset[n.Int64()]
	}
	return final, nil
}

func charsetFromChars(chars []Chars) []byte {
	var charset []byte
	for _, char := range chars {
		charset = append(charset, char.chars...)
	}
	return charset
}

func passesChars(generated data.Secret, chars []Chars) bool {
	for _, char := range chars {
		if !passesChar(generated, char) {
			return false
		}
	}
	return true
}

func passesChar(generated data.Secret, char Chars) bool {
	var total int
	for _, b := range char.chars {
		total += bytes.Count(generated, []byte{b})
		if total >= char.minRequired {
			return true
		}
	}
	return total >= char.minRequired
}
