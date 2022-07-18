package main_test

import (
	"fmt"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/error_handler"
	"github.com/stretchr/testify/assert"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestMainAdd(t *testing.T) {
	dir, err := os.Getwd()
	error_handler.HandlePanic(err)
	cmdPath := filepath.Join(dir, binName)

	adderTest := []struct {
		name            string
		commandExecuted *exec.Cmd
		expected        string
	}{
		{
			name:            "Given a one number",
			commandExecuted: exec.Command(cmdPath, "2"),
			expected:        "Sum of 2, equal 2 \n",
		},
		{
			name:            "Given a multiple number",
			commandExecuted: exec.Command(cmdPath, "2", "3", "5"),
			expected:        "Sum of 2,3,5, equal 10 \n",
		},
		{
			name:            "Given the --input-file data/input.txt should return the calculation from input.txt",
			commandExecuted: exec.Command(cmdPath, "--input-file", "data/input.txt"),
			expected:        "Sum of 4,5,32,100,867543, equal 867,684 \n",
		},
		{
			name:            "Given the few files should return the sum of the all numbers in files",
			commandExecuted: exec.Command(cmdPath, "--input-file", "data/input.txt", "--input-file", "data/input2.csv"),
			expected:        "Sum of 4,5,32,100,867543, equal 867,684 \n",
		},
		{
			name:            "Given no arguments should return calculation from data/input.txt file",
			commandExecuted: exec.Command(cmdPath),
			expected:        "Sum of 4,5,32,100,867543 equal 867,684 \n",
		},
	}

	for _, tt := range adderTest {
		t.Run(tt.name, func(t *testing.T) {
			sum := CommandLineOutput(tt.commandExecuted)
			assert.Equal(t, tt.expected, sum)
		})
	}
}

func CommandLineOutput(cmd *exec.Cmd) string {
	cmdStdIn, err := cmd.StdinPipe()
	error_handler.HandlePanic(err)

	err = cmdStdIn.Close()
	error_handler.HandlePanic(err)

	out, err := cmd.CombinedOutput()
	error_handler.HandlePanic(err)

	return string(out)
}

var (
	binName = "add"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")

	build := exec.Command("go", "build", "-o", binName)

	if err := build.Run(); err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
		error_handler.HandlePanic(err)
		os.Exit(1)
	}

	fmt.Println("Running tests....")
	result := m.Run()

	fmt.Println("Cleaning up...")
	err := os.Remove(binName)
	error_handler.HandlePanic(err)
	os.Exit(result)
}
