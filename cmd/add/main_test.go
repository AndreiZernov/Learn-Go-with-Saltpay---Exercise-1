package main_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

const binName = "add"

func TestMainAdd(t *testing.T) {
	dir, dirErr := os.Getwd()
	if dirErr != nil {
		t.Errorf("Cannot get current directory")
	}

	cmdPath := filepath.Join(dir, binName)

	adderTest := []struct {
		Name            string
		CommandExecuted *exec.Cmd
		Expected        string
	}{
		{
			Name:            "Given a one number",
			CommandExecuted: exec.Command(cmdPath, "2"),
			Expected:        "2\n",
		},
		{
			Name:            "Given a multiple number",
			CommandExecuted: exec.Command(cmdPath, "2", "3", "5"),
			Expected:        "10\n",
		},
		{
			Name:            "Given the --input-file data/input.txt should return the calculation from input.txt",
			CommandExecuted: exec.Command(cmdPath, "--input-file", "data/input.txt"),
			Expected:        "867,684\n",
		},
		{
			Name:            "Given the few files should return the sum of the all integers in files",
			CommandExecuted: exec.Command(cmdPath, "--input-file", "data/input.txt", "--input-file", "data/input2.csv"),
			Expected:        "867,684\n",
		},
		{
			Name:            "Given no arguments should return calculation from data/input.txt file",
			CommandExecuted: exec.Command(cmdPath),
			Expected:        "867,684\n",
		},
	}

	for _, tt := range adderTest {
		t.Run(tt.Name, func(t *testing.T) {
			sum := CommandLineOutput(t, tt.CommandExecuted)
			assert.Equal(t, tt.Expected, sum)
		})
	}
}

func CommandLineOutput(t testing.TB, cmd *exec.Cmd) string {
	t.Helper()
	cmdStdIn, err := cmd.StdinPipe()
	if err != nil {
		t.Errorf("Cannot create stdin pipe")
	}

	cmdStdIn.Close()

	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Cannot execute command")
	}

	return string(out)
}

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")

	build := exec.Command("go", "build", "-o", binName)

	if err := build.Run(); err != nil {
		log.Fatal("Cannot build tool")
	}

	fmt.Println("Running tests....")
	result := m.Run()

	fmt.Println("Cleaning up...")

	if err := os.Remove(binName); err != nil {
		log.Fatal("Cannot remove tool")
	}

	os.Exit(result)
}
