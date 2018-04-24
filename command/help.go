package command

import "fmt"

func HandleHelp(originalDbName string, args []string) {
	fmt.Println("Usage: pgbackup <command>")
	fmt.Println("  Commands:")
	for _, command := range Index {
		fmt.Printf("    %v - %v\n", command.DisplayCommand, command.Description)
	}
}
