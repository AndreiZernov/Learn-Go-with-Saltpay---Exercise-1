package main_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestCommandLine_GetArguments(t *testing.T) {
	t.Run("Given a one number", func(t *testing.T) {
		sum := cmdRunner("2")

		assert.Equal(t, "Sum of 2 equal 2 \n", sum)
	})

	t.Run("Given a one number", func(t *testing.T) {
		sum := cmdRunner("2, 3, 5")

		assert.Equal(t, "Sum of 2, 3, 5 equal 10 \n", sum)
	})

	t.Run("No arg passed", func(t *testing.T) {
		sum := cmdRunner("")

		assert.Equal(t, "Sum of 4, 5, 32, 100, 867543 equal 867,684 \n", sum)
	})
}

var (
	binName = "add"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")

	build := exec.Command("go", "build", "-o", binName)

	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
		os.Exit(1)
	}

	fmt.Println("Running tests....")
	result := m.Run()

	fmt.Println("Cleaning up...")
	os.Remove(binName)

	os.Exit(result)
}

func cmdRunner(arg string) string {
	var cmd *exec.Cmd

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	cmdPath := filepath.Join(dir, binName)

	if arg != "" {
		cmd = exec.Command(cmdPath, arg)
	} else {
		cmd = exec.Command(cmdPath)
	}

	cmdStdIn, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}

	cmdStdIn.Close()

	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	return string(out)
}
