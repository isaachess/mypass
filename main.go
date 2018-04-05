package main

import (
	"flag"
	"fmt"
	"mypass/cmd"
	"mypass/store"
	"os"
	"os/user"
	"path/filepath"
)

var (
	programName = "mypass"
	configPath  = flag.String("config-path", "", "The path to the config json file")
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func run() error {
	currentUser, err := user.Current()
	if err != nil {
		return err
	}

	defaultPath := filepath.Join(currentUser.HomeDir, fmt.Sprintf(".%s", programName))

	path := *configPath

	if path == "" {
		path = filepath.Join(defaultPath, "config.json")
	}

	config := LoadConfig(path, defaultPath)

	store := store.NewJSONStore(config.DataFileLocation)
	addCmd := cmd.NewCommand(cmd.AddName, cmd.AddFlags, cmd.NewAdd(store), nil)
	listCmd := cmd.NewCommand(cmd.ListName, cmd.ListFlags, cmd.NewList(store), nil)
	mainSubs := []*cmd.Command{
		addCmd,
		listCmd,
	}
	mainCmd := cmd.NewCommand(cmd.MainName, cmd.MainFlags, cmd.NewMain(), mainSubs)

	if err := store.Connect(); err != nil {
		return err
	}
	defer store.Close()

	mainCmd.Execute(os.Args[1:])

	return nil
}
