package main

import "math/rand"

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numbers = "1234567890"
const symbols = `()!@#$%^&*-+=|{}[]:;<>,.?\/`

const lettersNumbers = letters + numbers
const lettersNumbersSymbols = lettersNumbers + symbols

type passwords struct {
	pwLength int
}

func (p *passwords) Generate() []byte {
	b := make([]byte, p.pwLength)
	for i := range b {
		b[i] = lettersNumbers[rand.Intn(len(lettersNumbers))]
	}
	return b
}
