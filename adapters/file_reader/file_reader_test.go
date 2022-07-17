package file_reader_test

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/file_reader"
	"path/filepath"
	"runtime"
	"testing"
)

func TestReadFile(t *testing.T) {
	t.Run("Should read file located at data/input.txt", func(t *testing.T) {
		path := file_reader.GetFilePathname("data/input.txt")
		expected := "4\n5\n32\n100\n867543"
		got := file_reader.ReadFile(path)

		if got != expected {
			t.Errorf("got %q, want %q", got, expected)
		}
	})

	t.Run("Should read file located at data/input2.csv", func(t *testing.T) {
		path := file_reader.GetFilePathname("data/input2.csv")
		expected := "4,5,32,100,867543"
		got := file_reader.ReadFile(path)

		if got != expected {
			t.Errorf("got %q, want %q", got, expected)
		}
	})
}

var (
	_, b, _, _ = runtime.Caller(0)
	Root       = filepath.Join(filepath.Dir(b), "../..")
)

func TestGetFilePathname(t *testing.T) {
	t.Run("get file pathname data/input.txt in the file system", func(t *testing.T) {
		path := "/data/input.txt"
		expected := Root + path
		got := file_reader.GetFilePathname(path)

		if got != expected {
			t.Errorf("got %q, want %q", got, expected)
		}
	})

	t.Run("get file pathname data/input2.csv in the file system", func(t *testing.T) {
		path := "/data/input2.csv"
		expected := Root + path
		got := file_reader.GetFilePathname(path)

		if got != expected {
			t.Errorf("got %q, want %q", got, expected)
		}
	})
}
