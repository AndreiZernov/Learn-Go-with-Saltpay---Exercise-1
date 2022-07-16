package commandLine

import (
	"log"
	"os"
	"strings"
)

type CommandLine struct {
}

func New() *CommandLine {
	return &CommandLine{}
}

func (c CommandLine) GetArguments() string {
	toGetAllArgs := os.Args[1:]
	var numbers string

	if len(toGetAllArgs) == 0 {
		content, err := os.ReadFile("data/input.txt")
		if err != nil {
			log.Fatal(err)
		}
		numbers = strings.Replace(string(content), "\n", ", ", -1)
	} else {
		numbers = toGetAllArgs[0]
	}

	return numbers
}
