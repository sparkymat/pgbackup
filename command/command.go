package command

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Details struct {
	DisplayCommand string
	Description    string
	Handler        func(string, []string)
}

var Index map[string]Details

func init() {
	Index = map[string]Details{
		"help":    {"help", "Displays the help page", HandleHelp},
		"list":    {"list", "List the database backups", HandleList},
		"backup":  {"backup <tag>", "Backup the database with the given tag (or current timestamp if no tag provided)", HandleBackup},
		"restore": {"restore <tag>", "Restore from the tagged backup. Will auto-backup current database.", HandleRestore},
		"cleanup": {"cleanup", "Drops all auto backups", HandleCleanup},
	}
}

func displayError(originalDbName string, args []string) {
	fmt.Printf("Error: Unknown command\n\n")
	HandleHelp(originalDbName, args)
}

func HandleInput(originalDbName string, commandString string, args []string) {
	if commandDetails, exists := Index[commandString]; exists {
		commandDetails.Handler(originalDbName, args)
	} else {
		displayError(originalDbName, args)
	}
}

func execAndReturnOutput(commandString string) ([]string, error) {
	command := exec.Command("bash", "-c", commandString)
	command.Env = os.Environ()
	output, err := command.Output()
	if err != nil {
		return []string{}, errors.New("Unable to run command")
	}
	return strings.Split(string(output), "\n"), nil
}

func confirm() {
	var input string
	var err error

	fmt.Printf("Are you sure (yes/no)? ")
	input, err = bufio.NewReader(os.Stdin).ReadString('\n')
	if err == nil && (strings.TrimSpace(strings.ToLower(input)) == "yes" || strings.TrimSpace(strings.ToLower(input)) == "y") {
		return
	}

	panic("User abort. Exiting...")
}
