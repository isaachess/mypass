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
		subCo := c.findSub(subName)
		if subCo != nil {
			return subCo.Execute(args[1:])
		}
	}
	c.flags.Parse(args)
	if err := c.runner.Run(args); err != nil {
		fmt.Println(c.usage())
		return err
	}
	return nil
}

func (c *Command) usage() string {
	var usg = `
USAGE: %s

SUBCOMMANDS:
%s`
	return fmt.Sprintf(usg, c.runner.Usage(), c.subUsage())
}

func (c *Command) subUsage() string {
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
