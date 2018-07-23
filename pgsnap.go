package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/sparkymat/pgsnap/command"
)

type appConfig struct {
	DbName string `toml:"db_name"`
}

func main() {
	var configPath string
	var config appConfig
	var err error

	flag.StringVar(&configPath, "c", ".pgbackup.toml", "Path to the config file")
	flag.Parse()

	if _, err = toml.DecodeFile(configPath, &config); err != nil {
		fmt.Fprintf(os.Stderr, "Error: Unable to load config file. %v", err.Error())
		os.Exit(-1)
	}

	args := flag.Args()

	if len(args) == 0 {
		command.HandleHelp(config.DbName, []string{})
		os.Exit(0)
	}

	commandString := args[0]
	commandArgs := args[1:]

	command.HandleInput(config.DbName, commandString, commandArgs)
}
