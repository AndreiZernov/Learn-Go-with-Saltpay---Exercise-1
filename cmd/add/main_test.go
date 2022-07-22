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
		Name            string
		CommandExecuted *exec.Cmd
		Expected        string
	}{
		{
			Name:            "Given a one number",
			CommandExecuted: exec.Command(cmdPath, "2"),
			Expected:        "Sum of 2, equal 2 \n",
		},
		{
			Name:            "Given a multiple number",
			CommandExecuted: exec.Command(cmdPath, "2", "3", "5"),
			Expected:        "Sum of 2,3,5, equal 10 \n",
		},
		{
			Name:            "Given the --input-file data/input.txt should return the calculation from input.txt",
			CommandExecuted: exec.Command(cmdPath, "--input-file", "data/input.txt"),
			Expected:        "Sum of 4,5,32,100,867543, equal 867,684 \n",
		},
		{
			Name:            "Given the few files should return the sum of the all numbers in files",
			CommandExecuted: exec.Command(cmdPath, "--input-file", "data/input.txt", "--input-file", "data/input2.csv"),
			Expected:        "Sum of 4,5,32,100,867543, equal 867,684 \n",
		},
		{
			Name:            "Given no arguments should return calculation from data/input.txt file",
			CommandExecuted: exec.Command(cmdPath),
			Expected:        "Sum of 4,5,32,100,867543 equal 867,684 \n",
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
