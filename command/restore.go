package command

import (
	"fmt"
	"time"
)

func HandleRestore(originalDbName string, args []string) {
	if len(args) == 0 {
		panic("Unable to restore without a tag")
	}

	confirm()

	autobackupTag := restoreFromBackup(originalDbName, args[0])
	fmt.Printf("Successfully restored from %v_%v (previous one auto-backed up to %v_%v)\n", originalDbName, args[0], originalDbName, autobackupTag)
}

func restoreFromBackup(baseName string, tag string) string {
	autobackupTag := fmt.Sprintf("auto_%v", time.Now().Format("20060102150405"))

	autobackupName := fmt.Sprintf("%v_%v", baseName, autobackupTag)
	restoreName := fmt.Sprintf("%v_%v", baseName, tag)

	autobackupCommandString := fmt.Sprintf("psql -c \"ALTER DATABASE %v RENAME TO %v;\"", baseName, autobackupName)
	restoreCommandString := fmt.Sprintf("psql -c \"CREATE DATABASE %v TEMPLATE %v;\"", baseName, restoreName)

	if _, err := execAndReturnOutput(autobackupCommandString); err != nil {
		panic(fmt.Sprintf("Failed to autobackup database from %v to %v", baseName, autobackupName))
	}

	if _, err := execAndReturnOutput(restoreCommandString); err != nil {
		panic(fmt.Sprintf("Failed to restore database %v from template %v", baseName, restoreName))
	}

	return autobackupTag
}
