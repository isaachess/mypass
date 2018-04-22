package cmd

import (
	"errors"
	"flag"
	"fmt"
	"mypass/encrypt"
	"mypass/store"
	"mypass/terminal"

	"github.com/atotto/clipboard"
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

	return c.cpSiteName(args[0])
}

func (c *Cp) cpSiteName(name string) error {
	pw, err := c.store.Get(name)
	if err != nil && err != store.ErrorNotFound {
		return err
	}

	// TODO(isaac): way to alias a site for fast copying

	// check partial matches; if single partial match found, use it
	if err == store.ErrorNotFound {
		matches := c.store.MatchNames(name)
		if len(matches) != 1 {
			return c.findAndPrintPartials(name, matches)
		}
		return c.cpSiteName(matches[0])
	}

	mp, err := terminal.ReadMasterPassword()
	if err != nil {
		return err
	}

	key := encrypt.MasterToKey(mp, pw.Salt)

	decrypted, err := encrypt.Decrypt(pw.Hash, pw.Salt, key)
	if err != nil {
		return err
	}

	return clipboard.WriteAll(decrypted.DangerousString())
}

func (c *Cp) findAndPrintPartials(name string, matches []string) error {
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
