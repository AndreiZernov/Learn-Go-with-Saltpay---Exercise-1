package main_test

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"
)

func TestMainAdd(t *testing.T) {
	var cmd *exec.Cmd

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	cmdPath := filepath.Join(dir, binName)

	t.Run("Given --web-server in arguments should start the web server", func(t *testing.T) {
		cmd = exec.Command(cmdPath, "--web-server")

		cmdStdIn, err := cmd.StdinPipe()
		if err != nil {
			panic(err)
		}

		cmdStdIn.Close()

		waitForServer()
	})
}

func waitForServer() {
	for i := 0; i < 10; i++ {
		conn, _ := net.Dial("tcp", net.JoinHostPort("localhost", "8081"))
		if conn != nil {
			conn.Close()
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
}

var (
	binName = "math"
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
