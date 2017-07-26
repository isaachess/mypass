package main

import (
	"time"
)

type PasswordInfo struct {
	Name          string
	Salt          []byte
	Date          time.Time
	Length        int
	RequiredChars []*RequiredChars
}

type RequiredChars struct {
	Chars string
	Num   int
}

func NewPasswordInfo(name string) (*PasswordInfo, error) {
	return &PasswordInfo{
		Name:   name,
		Date:   time.Now(),
		Length: 12,
	}, nil
}

func (p *PasswordInfo) AddRequiredChars(rc *RequiredChars) {
	p.RequiredChars = append(p.RequiredChars, rc)
}
