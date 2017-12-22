package cmd

import "fmt"

type Command struct {
	parent string
	name   string
	subs   []*Command
}

func NewCommand(parent, name string) *Command {
	return &Command{parent: parent, name: name}
}

func (c *Command) AddSub(co *Command) {
	c.subs = append(c.subs, co)
}

func (c *Command) Usage() string {
	return fmt.Sprintf(`
	USAGE: %s
	SUBCOMMANDS:

	%s
	`, c.usage(), c.subCommandUsage())
}

func (c *Command) usage() string {
	var parent = c.parent
	if parent != "" {
		parent = parent + " "
	}
	return fmt.Sprintf("%s%s", parent, c.name)
}

func (c *Command) subCommandUsage() string {
	if len(c.subs) < 1 {
		return ""
	}

	var final []byte
	for _, sub := range c.subs {
		final = append(final, '\n')
		final = append(final, []byte(sub.name)...)
	}
	return string(final)
}
