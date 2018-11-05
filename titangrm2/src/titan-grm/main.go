package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"

	"titan-grm/command"
)

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
	return runCustom(args, command.MakeCommands(nil))
}

func runCustom(args []string, commands map[string]cli.CommandFactory) int {
	for _, arg := range args {
		if arg == "-v" || arg == "-version" || arg == "--version" {
			args = []string{"version"}
			break
		}
	}

	// Build the commands to include in the help now.
	names := make([]string, 0, len(commands))
	for c := range commands {
		names = append(names, c)
	}

	cli := &cli.CLI{
		Args:     args,
		Commands: commands,
		HelpFunc: cli.FilteredHelpFunc(names, cli.BasicHelpFunc("titan-grm")),
	}

	exitCode, err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
		return 1
	}
	return exitCode
}
