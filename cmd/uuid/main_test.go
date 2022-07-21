package main_test

import (
	"fmt"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/error_handler"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/files"
	"github.com/stretchr/testify/assert"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestMainUUID(t *testing.T) {
	dir, err := os.Getwd()
	error_handler.HandlePanic(err)
	cmdPath := filepath.Join(dir, binName)
	pathname := "test_authorised_api_access_keys.txt"
	t.Setenv("AUTH_KEYS_PATHNAME", "test_authorised_api_access_keys.txt")

	t.Run("Given a number 1 should generate a test_authorised_api_access_keys file with one UUID", func(t *testing.T) {
		out := CommandLineOutput(t, exec.Command(cmdPath, "2"))

		data := files.ReadFile(pathname)
		files.RemoveFile(pathname)

		assert.Equal(t, 74, len(data))
		assert.Equal(t, "Successfully generated 2 uuid keys in test_authorised_api_access_keys.txt \n", out)
	})

	t.Run("By not specifying a number should not generate new keys", func(t *testing.T) {
		out := CommandLineOutput(t, exec.Command(cmdPath))

		assert.Equal(t, "Keys was not generated, please specify the amount of keys to generate (go run uuid 1000) \n", out)
	})
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
	binName = "uuid"
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
