package main_test

import (
	"fmt"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/files"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

const (
	binName              = "uuid"
	envAuthKeysEnvName   = "AUTH_KEYS_PATHNAME"
	testAuthKeysPathname = "test_authorised_api_access_keys.txt"
)

func TestMainUUID(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Errorf("Cannot get current directory")
	}

	cmdPath := filepath.Join(dir, binName)
	t.Setenv(envAuthKeysEnvName, testAuthKeysPathname)

	t.Run("Given a number 1 should generate a test_authorised_api_access_keys file with one UUID", func(t *testing.T) {
		out := CommandLineOutput(t, exec.Command(cmdPath, "2"))

		data, err := files.ReadFile(testAuthKeysPathname)
		if err != nil {
			t.Errorf("Cannot read file %s", testAuthKeysPathname)
		}
		files.RemoveFile(testAuthKeysPathname)

		assert.Equal(t, 74, len(data))
		assert.Equal(t, "Successfully generated 2 uuid keys \nTo generate 2 keys it took 0 Seconds \n", out)
	})

	t.Run("By not specifying a number should not generate new keys", func(t *testing.T) {
		out := CommandLineOutput(t, exec.Command(cmdPath))

		assert.Equal(t, "Keys was not generated, please specify the amount of keys to generate (go run uuid 1000) \n", out)
	})
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
		t.Errorf("Cannot run command")
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
