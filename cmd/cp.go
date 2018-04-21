package cmd

import (
	"errors"
	"flag"
	"fmt"
	"mypass/store"
	"mypass/terminal"
)

const CpName = "cp"

var CpFlags = flag.NewFlagSet(CpName, flag.ExitOnError)

type Cp struct {
	store store.Store
}

func NewCp(store store.Store) *Cp {
	return &Cp{store: store}
}

func (c *Cp) Run(args []string) error {
	if len(args) != 1 {
		return errors.New("Please provide which password to cp")
	}

	site_name := args[0]

	pw, err := c.store.Get(site_name)
	if err != nil && err != store.ErrorNotFound {
		return err
	}

	// TODO(isaac): If partials come back with *one* match, we should just cp
	// that to the keyboard and print a message of which site we printed

	// TODO(isaac): way to alias a site for fast copying
	if err == store.ErrorNotFound {
		return c.findAndPrintPartials(site_name)
	}

	mp, err := terminal.ReadMasterPassword()
	if err != nil {
		return err
	}

	fmt.Println("pw", pw, "mp", mp)

	return nil

	//key := encrypt.MasterToKey(mp, salt)

	//hash, err := encrypt.Encrypt(pw, salt, key)
	//if err != nil {
	//return err
	//}

	//pi := data.NewPasswordInfo(string(username), hash, salt)
	//return a.store.Put(string(name), pi)
}

func (c *Cp) findAndPrintPartials(name string) error {
	matches := c.store.MatchNames(name)
	if len(matches) == 0 {
		fmt.Println("No matches found for name: ", name)
		return nil
	}

	var matchesFormat string
	for _, match := range matches {
		matchesFormat += fmt.Sprintf("%s\n", match)
	}

	var matchesMsg = `
No password found for name "%s". Did you mean:

%s
`
	fmt.Printf(matchesMsg, name, matchesFormat)
	return nil
}

func (c *Cp) Usage() string {
	return fmt.Sprintf("%s [site-name]", CpName)
}
