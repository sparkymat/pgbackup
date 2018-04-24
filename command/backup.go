package command

import (
	"fmt"
	"regexp"
	"time"
)

func HandleBackup(originalDbName string, args []string) {
	var tag string

	if len(args) == 0 {
		tag = time.Now().Format("20060102150405")
	} else {
		tag = args[0]

		if match, _ := regexp.MatchString("^[a-z0-9]+$", tag); !match {
			panic("Invalid tag passed to 'backup'")
		}
	}

	backupDatabase(originalDbName, tag)
	fmt.Printf("Successfully backed up %v to %v_%v\n", originalDbName, originalDbName, tag)
}

func backupDatabase(baseName string, tagName string) {
	backupName := fmt.Sprintf("%v_%v", baseName, tagName)

	backupCommandString := fmt.Sprintf("psql -c \"ALTER DATABASE %v RENAME TO %v;\"", baseName, backupName)
	copyCommandString := fmt.Sprintf("psql -c \"CREATE DATABASE %v TEMPLATE %v;\"", baseName, backupName)

	if _, err := execAndReturnOutput(backupCommandString); err != nil {
		panic(fmt.Sprintf("Failed to rename database from %v to %v", baseName, backupName))
	}

	if _, err := execAndReturnOutput(copyCommandString); err != nil {
		panic(fmt.Sprintf("Failed to create new database %v from template %v", baseName, backupName))
	}
}
