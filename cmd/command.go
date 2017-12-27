package cmd

import (
	"flag"
	"fmt"
)

type Executor interface {
	Run(args []string) error
}

type Command struct {
	exec  Executor
	flags *flag.FlagSet
	name  string
	subs  []*Command
}

func NewCommand(name string, flags *flag.FlagSet, exec Executor, subs []*Command) *Command {
	return &Command{
		exec:  exec,
		flags: flags,
		name:  name,
		subs:  subs,
	}
}

func (c *Command) Execute(args []string) error {
	// Check for a sub, if it finds one return that execute
	// If no sub found, execute this one
	if len(args) > 0 {
		subName := args[0]
		subCo := c.findSub(subName)
		if subCo != nil {
			return subCo.Execute(args[1:])
		}
	}
	c.flags.Parse(args)
	if err := c.exec.Run(args); err != nil {
		fmt.Println(c.usage())
		return err
	}
	return nil
}

func (c *Command) usage() string {
	var usg = `
USAGE: %s [sub-command] [args]

SUBCOMMANDS:
%s`
	return fmt.Sprintf(usg, c.name, c.subUsage())
}

func (c *Command) subUsage() string {
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
