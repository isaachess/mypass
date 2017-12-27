package data

import "time"

type PasswordInfo struct {
	Date     time.Time `json:"date"`
	Hash     []byte    `json:"hash"`
	Salt     []byte    `json:"salt"`
	Username []byte    `json:"username"`
}

func NewPasswordInfo(username, hash, salt []byte) *PasswordInfo {
	return &PasswordInfo{
		Date:     time.Now(),
		Hash:     hash,
		Salt:     salt,
		Username: username,
	}
}
