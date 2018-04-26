package cmd

import (
	"flag"
	"fmt"
)

type Runner interface {
	Run(args []string) error
	Usage() string
}

type Command struct {
	runner Runner
	flags  *flag.FlagSet
	name   string
	subs   []*Command
}

func NewCommand(name string, flags *flag.FlagSet, runner Runner,
	subs []*Command) *Command {
	return &Command{
		runner: runner,
		flags:  flags,
		name:   name,
		subs:   subs,
	}
}

func (c *Command) Execute(args []string) error {
	// Check for a sub, if it finds one return that execute
	// If no sub found, execute this one
	if len(args) > 0 {
		subName := args[0]
		subCom := c.findSub(subName)
		if subCom != nil {
			return subCom.Execute(args[1:])
		}
	}

	// if no matching sub found but this command has subs, print usage
	if len(c.subs) > 0 {
		fmt.Println(c.subUsage())
		return nil
	}

	c.flags.Parse(args)
	if err := c.runner.Run(args); err != nil {
		fmt.Println(c.errorText(err))
		return err
	}
	return nil
}

func (c *Command) subUsage() string {
	var usg = `
Usage: %s [sub-command]

Subcommands:

%s`
	return fmt.Sprintf(usg, c.name, c.subCommands())
}

func (c *Command) errorText(err error) string {
	var msg = `
The following error occured: %s

%s`
	return fmt.Sprintf(msg, err.Error(), c.usage())
}

func (c *Command) usage() string {
	var usg = "Usage: %s"
	return fmt.Sprintf(usg, c.runner.Usage())
}

func (c *Command) subCommands() string {
	if len(c.subs) == 0 {
		return "none"
	}

	var final string
	for _, sub := range c.subs {
		final += fmt.Sprintf("%s\n", sub.name)
	}
	return final
}

func (c *Command) findSub(subName string) *Command {
	for _, co := range c.subs {
		if co.name == subName {
			return co
		}
	}
	return nil
}
