package terminal

import (
	"bufio"
	"fmt"
	"mypass/data"
	"os"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func ReadLine(prompt string) (line []byte, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	line, _, err = reader.ReadLine()
	return line, err
}

func ReadMasterPassword() (data.Secret, error) {
	return readPassword("Master password: ")
}

func ConfirmMasterPassword() (data.Secret, error) {
	return readPassword("Confirm password: ")
}

func ReadPassword() (data.Secret, error) {
	return readPassword("Password: ")
}

func readPassword(prompt string) (data.Secret, error) {
	fmt.Print(prompt)
	pw, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return nil, err
	}
	fmt.Println("") // this ensures a new line occurs after enter on password
	return data.NewSecret(pw), nil
}
