package cmd

import (
	"flag"
	"mypass/data"
	"mypass/encrypt"
	"mypass/generate"
	"mypass/store"
	"mypass/terminal"
)

const AddName = "add"

var AddFlags = flag.NewFlagSet(AddName, flag.ExitOnError)

type Add struct {
	store *store.JSONStore
}

func NewAdd(store *store.JSONStore) *Add {
	return &Add{store: store}
}

func (a *Add) Run(args []string) error {
	mp, err := terminal.ReadMasterPassword()
	if err != nil {
		return err
	}

	name, err := terminal.ReadLine("Site Name: ")
	if err != nil {
		return err
	}

	username, err := terminal.ReadLine("Username: ")
	if err != nil {
		return err
	}

	pw, err := generate.Generate(12, generate.StrictChars)
	if err != nil {
		return err
	}

	salt, err := encrypt.NewSalt(12)
	if err != nil {
		return err
	}

	key := encrypt.MasterToKey(mp, salt)

	hash, err := encrypt.Encrypt(pw, salt, key)
	if err != nil {
		return err
	}

	pi := data.NewPasswordInfo(string(username), hash, salt)
	return a.store.Put(string(name), pi)
}
