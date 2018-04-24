package command

import (
	"fmt"
	"strings"
)

func HandleCleanup(originalDbName string, args []string) {
	confirm()

	backups := listBackups(originalDbName)

	for _, backup := range backups {
		if strings.HasPrefix(backup, fmt.Sprintf("%v_auto_", originalDbName)) {
			cleanup(backup)
		}
	}
}

func cleanup(databaseName string) {
	commandString := fmt.Sprintf("psql -c \"DROP DATABASE %v;\"", databaseName)
	if _, err := execAndReturnOutput(commandString); err != nil {
		panic(fmt.Sprintf("Failed to drop database %v", databaseName))
	}

	fmt.Printf("Dropped database %v\n", databaseName)
}
