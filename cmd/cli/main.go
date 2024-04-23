package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/paisit04/building-modern-cli/cmd/cli/command"
	"github.com/paisit04/building-modern-cli/internal/interfaces"
)

func main() {
	client := &http.Client{}
	cmds := []interfaces.Command{
		command.NewRandomCommand(client),
		command.NewGetCommand(client),
		command.NewUploadCommand(client),
		// command.newListCommand(client),
	}
	parser := command.NewParser(cmds)
	if err := parser.Parse(os.Args[1:]); err != nil {
		os.Stderr.WriteString(fmt.Sprintf("error: %v", err.Error()))
		os.Exit(1)
	}
}
