package main_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestMainAdd(t *testing.T) {
	var cmd *exec.Cmd

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	cmdPath := filepath.Join(dir, binName)

	t.Run("Given a one number", func(t *testing.T) {
		cmd = exec.Command(cmdPath, "2")

		cmdStdIn, err := cmd.StdinPipe()
		if err != nil {
			panic(err)
		}

		cmdStdIn.Close()

		out, err := cmd.CombinedOutput()
		if err != nil {
			panic(err)
		}

		sum := string(out)

		assert.Equal(t, "Sum of 2, equal 2 \n", sum)
	})

	t.Run("Given a multiple number", func(t *testing.T) {
		cmd = exec.Command(cmdPath, "2", "3", "5")

		cmdStdIn, err := cmd.StdinPipe()
		if err != nil {
			panic(err)
		}

		cmdStdIn.Close()

		out, err := cmd.CombinedOutput()
		if err != nil {
			panic(err)
		}

		sum := string(out)

		assert.Equal(t, "Sum of 2,3,5, equal 10 \n", sum)
	})

	t.Run("No arg passed", func(t *testing.T) {
		cmd = exec.Command(cmdPath)

		cmdStdIn, err := cmd.StdinPipe()
		if err != nil {
			panic(err)
		}

		cmdStdIn.Close()

		out, err := cmd.CombinedOutput()
		if err != nil {
			panic(err)
		}

		sum := string(out)

		assert.Equal(t, "Sum of 4,5,32,100,867543 equal 867,684 \n", sum)
	})

	t.Run("Given the --input-file data/input.txt should return the calculation of numbers inside the file input.txt", func(t *testing.T) {
		cmd = exec.Command(cmdPath, "--input-file", "data/input.txt")

		cmdStdIn, err := cmd.StdinPipe()
		if err != nil {
			panic(err)
		}

		cmdStdIn.Close()

		out, err := cmd.CombinedOutput()
		if err != nil {
			panic(err)
		}

		sum := string(out)

		assert.Equal(t, "Sum of 4,5,32,100,867543, equal 867,684 \n", sum)
	})

	t.Run("Given the few files should return the sum of the all files", func(t *testing.T) {
		cmd = exec.Command(cmdPath, "--input-file", "data/input.txt", "--input-file", "data/input2.csv")

		cmdStdIn, err := cmd.StdinPipe()
		if err != nil {
			panic(err)
		}

		cmdStdIn.Close()

		out, err := cmd.CombinedOutput()
		if err != nil {
			panic(err)
		}

		sum := string(out)

		assert.Equal(t, "Sum of 4,5,32,100,867543, equal 867,684 \n", sum)
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
