package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	version = "1.0.0"

	help = `tagline <subcommand> <file>
 help - Show this message.
 version - Show the current version.
 read - Read and print the metadata of an audio file.
 write - Write the metadata of an audio file.`
)

func _main() error {
	args := os.Args[1:]

	if len(args) < 1 {
		return errors.New("Expected 'tagline <subcommand>'")
	}

	subcmd := args[0]

	var subfunc func(string, []string) error

	switch subcmd {
	case "help":
		fmt.Println(help)
		return nil
	case "version":
		fmt.Println("Tagline version", version)
		return nil
	case "read":
		subfunc = subRead
	case "write":
		subfunc = subWrite
	default:
		return fmt.Errorf("Unknown subcommand '%s'", subcmd)
	}

	if len(args) < 2 {
		return fmt.Errorf("Expected 'tagline %s <file>'", subcmd)
	}

	if args[1] == "-h" || args[1] == "--help" {
		return subfunc("", args[1:])
	} else if strings.HasPrefix(args[1], "-") {
		return fmt.Errorf("Expected 'tagline %s <file>'", subcmd)
	} else {
		return subfunc(args[1], args[2:])
	}
}

func main() {
	if err := _main(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
