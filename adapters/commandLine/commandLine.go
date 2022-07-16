package commandLine

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type CommandLine struct {
}

func New() *CommandLine {
	return &CommandLine{}
}

var (
	_, b, _, _ = runtime.Caller(0)
	Root       = filepath.Join(filepath.Dir(b), "../..")
	numbers    string
)

func (c CommandLine) GetArguments() string {
	toGetAllArgs := os.Args[1:]

	switch {
	case len(toGetAllArgs) == 0:
		filepath := filepath.Join(Root, "/data/", "input.txt")
		content, err := ioutil.ReadFile(filepath)
		if err != nil {
			log.Fatal(err)
		}
		numbers = strings.Replace(string(content), "\n", ", ", -1)
	case toGetAllArgs[0] == "--input-file":
		filepath := filepath.Join(Root, "/", toGetAllArgs[1])
		content, err := ioutil.ReadFile(filepath)
		if err != nil {
			log.Fatal(err)
		}
		numbers = strings.Replace(string(content), "\n", ", ", -1)
	default:
		numbers = toGetAllArgs[0]
	}

	return numbers
}
