package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func HandleList(originalDbName string, args []string) {
	backups := listBackups(originalDbName)

	prefixPrint := color.New(color.Italic, color.FgWhite).FprintfFunc()
	tagPrint := color.New(color.Bold, color.FgHiWhite).FprintfFunc()

	fmt.Printf("Found %d backup(s):\n", len(backups))
	for _, backupName := range backups {
		prefixPrint(os.Stdout, "[%v_]", originalDbName)
		tagPrint(os.Stdout, strings.Replace(backupName, fmt.Sprintf("%v_", originalDbName), "", -1))
		color.New(color.FgWhite).Println("")
	}
}

func listBackups(originalDbName string) []string {
	var outputLines []string
	var err error

	if outputLines, err = execAndReturnOutput("psql -c \"\\l\""); err != nil {
		panic(fmt.Sprintf("'list' command failed. Error: %v", err.Error()))
	}

	matchedNames := []string{}
	for _, eachLine := range outputLines {
		if strings.Contains(eachLine, fmt.Sprintf("%v_", originalDbName)) {
			words := strings.Split(eachLine, "|")
			matchedNames = append(matchedNames, strings.Trim(words[0], " "))
		}
	}

	return matchedNames
}
