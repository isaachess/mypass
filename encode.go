package main

import "encoding/base64"

const lettersNumbers = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789078"
const letNumSpec = "abcdefghijklmnpqrstuvwxyzABCDEFGHIJKLMNPQRSTUVWXYZ1234567890!&@#"

var lettersNumbersEnc = base64.NewEncoding(lettersNumbers).WithPadding(base64.NoPadding)
var letNumSpecEnc = base64.NewEncoding(letNumSpec).WithPadding(base64.NoPadding)

func EncodeLettersNumbers(d []byte) string {
	return lettersNumbersEnc.EncodeToString(d)
}

func EncodeLettersNumbersSpecial(d []byte) string {
	return letNumSpecEnc.EncodeToString(d)
}
