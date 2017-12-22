package data

import "time"

type PasswordInfo struct {
	Salt []byte    `json:"salt"`
	Hash []byte    `json:"hash"`
	Date time.Time `json:"date"`
}

func NewPasswordInfo(hash, salt []byte) *PasswordInfo {
	return &PasswordInfo{
		Date: time.Now(),
		Salt: salt,
		Hash: hash,
	}
}
