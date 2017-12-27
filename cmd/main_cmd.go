package cmd

import (
	"errors"
	"flag"
)

const MainName = "mypass"

var MainFlags = flag.NewFlagSet(MainName, flag.ExitOnError)

type Main struct{}

func NewMain() *Main { return &Main{} }

func (m *Main) Run(args []string) error {
	return errors.New("Main should never be called on its own")
}
