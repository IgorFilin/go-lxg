package main

import (
	"mycli/cmd/commands"
	"mycli/cmd/root"
)

func main() {
	commands.StartExecute()
	root.Execute()
}