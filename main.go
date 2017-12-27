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
	mainCmd := cmd.NewCommand(cmd.MainName, cmd.MainFlags, cmd.NewMain(), []*cmd.Command{addCmd})

	if err := store.Connect(); err != nil {
		return err
	}
	defer store.Close()

	mainCmd.Execute(os.Args[1:])

	return nil

	//master := []byte("sample master passwd")
	//pw, err := generate.Generate(8, generate.AllChars)
	//if err != nil {
	//return err
	//}

	//salt, err := encrypt.NewSalt(12)
	//if err != nil {
	//return err
	//}

	//key := encrypt.MasterToKey(master, salt)

	//hash, err := encrypt.Encrypt(pw, salt, key)
	//if err != nil {
	//return err
	//}

	//b := data.NewEncodableBytes(hash, base64.StdEncoding)

	//fmt.Printf("encrypted %s to %s\n", string(pw), b)

	//str := b.String()
	//decoded, err := base64.StdEncoding.DecodeString(str)
	//if err != nil {
	//return err
	//}

	//dec, err := encrypt.Decrypt(decoded, salt, key)
	//if err != nil {
	//return err
	//}

	//fmt.Printf("decrypted %s\n", string(dec))
	//return nil
}
