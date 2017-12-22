package main

import (
	"errors"
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
	addCmd := flag.NewFlagSet(cmd.AddCmd, flag.ExitOnError)

	if len(os.Args) < 2 {
		return errors.New("Please check usage")
	}

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

	store := store.NewJSONStore(config.dataFileLocation)

	fmt.Println("loaded this config", config)
	if err := store.Connect(); err != nil {
		return err
	}
	defer store.Close()

	switch os.Args[1] {
	case cmd.AddCmd:
		addCmd.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		return nil
	}

	if addCmd.Parsed() {
		return cmd.NewAdd(store).Run()
	}

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
