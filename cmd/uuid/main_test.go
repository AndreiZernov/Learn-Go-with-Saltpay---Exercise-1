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

const binName = "uuid"
const pathname = "test_authorised_api_access_keys.txt"

func TestMainUUID(t *testing.T) {
	dir, dirErr := os.Getwd()
	error_handler.AnnotatingError(dirErr, "Cannot get current directory")
	cmdPath := filepath.Join(dir, binName)
	t.Setenv("AUTH_KEYS_PATHNAME", pathname)

	t.Run("Given a number 1 should generate a test_authorised_api_access_keys file with one UUID", func(t *testing.T) {
		out := CommandLineOutput(t, exec.Command(cmdPath, "2"))

		data := files.ReadFile(pathname)
		files.RemoveFile(pathname)

		assert.Equal(t, 74, len(data))
		assert.Equal(t, "Successfully generated 2 uuid keys in test_authorised_api_access_keys.txt \nTo generate 2 keys it took 0 Seconds \n", out)
	})

	t.Run("By not specifying a number should not generate new keys", func(t *testing.T) {
		out := CommandLineOutput(t, exec.Command(cmdPath))

		assert.Equal(t, "Keys was not generated, please specify the amount of keys to generate (go run uuid 1000) \n", out)
	})
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
