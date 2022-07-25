package main_test

import (
	"fmt"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/files"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/internals/testing_helpers"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

const (
	binName               = "math"
	envAuthKeysEnvName    = "AUTH_KEYS_PATHNAME"
	envLogPathname        = "LOG_PATHNAME"
	testAuthKeysPathname  = "test_authorised_api_access_keys.txt"
	testAccessLogPathname = "adapters/files/test_access_log.txt"
)

func TestMainFibo(t *testing.T) {
	t.Skip("Skipping test")

	t.Setenv(envAuthKeysEnvName, testAuthKeysPathname)
	t.Setenv(envLogPathname, testAccessLogPathname)

	defer files.RemoveFile(testAuthKeysPathname)
	defer files.RemoveFile(testAccessLogPathname)

	t.Run("Should successfully start the web server and return response", func(t *testing.T) {
		files.UUIDGenerator(1)

		dir, err := os.Getwd()
		if err != nil {
			t.Errorf("Cannot get current directory")
		}

		cmdServer := filepath.Join(dir, "math")
		exec.Command(cmdServer, "--web-server")

		out := testing_helpers.CaptureOutput(func() {

			cmdPath := filepath.Join(dir, binName)
			exec.Command(cmdPath, "2")
		})

		assert.Equal(t, "2", strings.TrimSpace(out))
	})
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
