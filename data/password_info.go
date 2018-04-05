package data

import (
	"fmt"
	"time"
)

type PasswordInfo struct {
	Date     time.Time `json:"date"`
	Hash     []byte    `json:"hash"`
	Salt     []byte    `json:"salt"`
	Username string    `json:"username"`
}

func NewPasswordInfo(username string, hash, salt []byte) *PasswordInfo {
	return &PasswordInfo{
		Date:     time.Now(),
		Hash:     hash,
		Salt:     salt,
		Username: username,
	}
}

func (p *PasswordInfo) String() string {
	return fmt.Sprintf(`Username: %s
Date Added: %s`, p.Username, p.Date)
}
