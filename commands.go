package main

import (
	"bufio"
	"fmt"
	"os"
)

type pwder interface {
	Generate() []byte
	Encrypt([]byte) ([]byte, error)
	Decrypt(key string, enc []byte) (string, error)
}

type store interface {
	Insert(key string, enc []byte) error
	Get(key string) ([]byte, error)
}

type commands struct {
	pwder pwder
	store store
}

func (c *commands) Generate() error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Password lookup name: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	pwd := c.pwder.Generate()

	enc, err := c.pwder.Encrypt(pwd)
	if err != nil {
		return err
	}

	if err := c.store.Insert(text, enc); err != nil {
		return err
	}

	fmt.Println("Generated password:", pwd)

	return nil
}

func (c *commands) Copy() error {
	return nil
}
