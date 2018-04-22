package cmd

import (
	"errors"
	"flag"
	"fmt"
	"mypass/data"
	"mypass/encrypt"
	"mypass/generate"
	"mypass/store"
	"mypass/terminal"
	"reflect"
)

const AddName = "add"

var AddFlags = flag.NewFlagSet(AddName, flag.ExitOnError)

type Add struct {
	store store.Store
}

func NewAdd(store store.Store) *Add {
	return &Add{store: store}
}

func (a *Add) Run(args []string) error {
	if len(args) > 0 {
		return errors.New("No args expected")
	}

	mp, err := terminal.ReadMasterPassword()
	if err != nil {
		return err
	}

	names, err := a.store.GetNames()
	if err != nil {
		return err
	}

	if len(names) == 0 {
		// this is the first pw, confirm mp
		mp2, err := terminal.ConfirmMasterPassword()
		if err != nil {
			return err
		}

		if !reflect.DeepEqual(mp, mp2) {
			fmt.Println("Master passwords do not match")
			return nil
		}
	} else {
		// this is not the first pw, make sure master password is the same used
		// for other sites
		ok, err := a.validateMasterPassword(mp)
		if err != nil {
			return err
		}
		if !ok {
			fmt.Println("Master password does not match previous passwords")
			return nil
		}
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

func (a *Add) validateMasterPassword(mp data.Secret) (bool, error) {
	// validate master password by seeing if we can use it to decrypt another
	// password already stored
	names, err := a.store.GetNames()
	if err != nil {
		return false, err
	}

	pw, err := a.store.Get(names[0])
	if err != nil {
		return false, err
	}

	key := encrypt.MasterToKey(mp, pw.Salt)

	_, err = encrypt.Decrypt(pw.Hash, pw.Salt, key)
	return err == nil, nil
}

func (a *Add) Usage() string {
	return AddName
}
