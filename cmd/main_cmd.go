package cmd

import (
	"errors"
	"flag"
	"fmt"
)

const MainName = "mypass"

var MainFlags = flag.NewFlagSet(MainName, flag.ExitOnError)

type Main struct{}

func NewMain() *Main { return &Main{} }

func (m *Main) Run(args []string) error {
	return errors.New("Missing sub-command")
}

func (m *Main) Usage() string {
	return fmt.Sprintf("%s [sub-command]", MainName)
}
