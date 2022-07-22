package main_test

import (
	"fmt"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/files"
	router "github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/temphttp"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

const binName = "math"
const envAuthKeysEnvName = "AUTH_KEYS_PATHNAME"
const envLogPathname = "LOG_PATHNAME"
const testAuthKeysPathname = "test_authorised_api_access_keys.txt"
const testAccessLogPathname = "adapters/files/test_access_log.txt"

func TestMainMath(t *testing.T) {
	t.Setenv(envAuthKeysEnvName, testAuthKeysPathname)
	t.Setenv(envLogPathname, testAccessLogPathname)

	defer files.RemoveFile(testAuthKeysPathname)
	defer files.RemoveFile(testAccessLogPathname)

	t.Run("Should return an error message about the missing the auth keys file", func(t *testing.T) {
		err := files.FindFile(testAuthKeysPathname)
		if err == nil {
			files.RemoveFile(testAuthKeysPathname)
		}

		dir, err := os.Getwd()
		if err != nil {
			t.Fatal(err)
		}

		cmdPath := filepath.Join(dir, binName)
		cmd := exec.Command(cmdPath)
		out, _ := cmd.CombinedOutput()

		assert.Equal(t, "Keys was not generated yet, please run the command to generate auth keys \x1b[0;32m go run cmd/uuid/main.go 1000 \x1b[0m and try it again", string(out))
	})

	t.Run("Should return an error message about the missing argument in command", func(t *testing.T) {
		files.UUIDGenerator(1)

		dir, err := os.Getwd()
		if err != nil {
			t.Fatal(err)
		}

		cmdPath := filepath.Join(dir, binName)
		cmd := exec.Command(cmdPath)
		out, _ := cmd.CombinedOutput()

		assert.Equal(t, "Web server did not start. Please check the command, should contain --web-server \n", string(out))
	})

	t.Run("Should successfully start the web server and return response", func(t *testing.T) {
		files.UUIDGenerator(1)

		dir, err := os.Getwd()
		if err != nil {
			t.Fatal(err)
		}

		cmdPath := filepath.Join(dir, binName)
		exec.Command(cmdPath, "--web-server")

		authKeys, err := files.ReadFile(testAuthKeysPathname)
		if err != nil {
			t.Fatal(err)
		}
		authKey := strings.Split(authKeys, "\n")[0]

		request, _ := http.NewRequest(http.MethodPost, "/add?num=2", nil)
		request.Header.Set("Authorization", authKey)
		response := httptest.NewRecorder()

		router := router.NewRouter()
		router.ServeHTTP(response, request)

		gotBody := response.Body.String()
		gotCode := response.Code

		assert.Equal(t, "Sum of 2 equal 2 \n", gotBody)
		assert.Equal(t, 200, gotCode)
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
