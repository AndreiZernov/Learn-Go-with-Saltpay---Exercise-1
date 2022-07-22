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

const binName = "add"

func TestMainAdd(t *testing.T) {
	dir, dirErr := os.Getwd()
	error_handler.AnnotatingError(dirErr, "Cannot get current directory")
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
	cmdStdIn, createErr := cmd.StdinPipe()
	error_handler.AnnotatingError(createErr, "Cannot create stdin pipe")

	closeErr := cmdStdIn.Close()
	error_handler.AnnotatingError(closeErr, "Cannot close stdin pipe")

	out, executeErr := cmd.CombinedOutput()
	error_handler.AnnotatingError(executeErr, "Cannot execute command")

	return string(out)
}

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")

	build := exec.Command("go", "build", "-o", binName)

	if buildErr := build.Run(); buildErr != nil {
		error_handler.AnnotatingError(buildErr, "Cannot build tool")
		os.Exit(1)
	}

	fmt.Println("Running tests....")
	result := m.Run()

	fmt.Println("Cleaning up...")
	removeToolErr := os.Remove(binName)
	error_handler.AnnotatingError(removeToolErr, "Cannot remove tool")
	os.Exit(result)
}
