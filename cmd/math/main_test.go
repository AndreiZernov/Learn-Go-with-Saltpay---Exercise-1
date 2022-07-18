package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

//func router() http.Handler {
//	mux := http.NewServeMux()
//	mux.HandleFunc("/", nil)
//	return mux
//}

//func TestMainAdd(t *testing.T) {
//	router := router()
//	ts := httptest.NewServer(router)
//	defer ts.Close()
//
//	dir, err := os.Getwd()
//	if err != nil {
//		panic(err)
//	}
//	cmdPath := filepath.Join(dir, binName)
//	cmd := exec.Command(cmdPath, "--web-server")
//	assert.NoError(t, cmd.Start())
//
//	newreq := func(method, url string, body io.Reader) *http.Request {
//		r, err := http.NewRequest(method, url, body)
//		if err != nil {
//			t.Fatal(err)
//		}
//		return r
//	}
//
//	tests := []struct {
//		name string
//		r    *http.Request
//	}{
//		{name: "2: testing post", r: newreq("POST", ts.URL+"/add?num=2&num=3", nil)},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			resp, err := http.DefaultClient.Do(tt.r)
//			body, err := io.ReadAll(resp.Body)
//			if err != nil {
//				t.Fatal(err)
//			}
//			assert.Equal(t, "Sum of 2,3, equal 5 \n", string(body))
//		})
//	}
//}

var (
	binName = "math"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")

	build := exec.Command("go", "build", "-o", binName)

	if err := build.Run(); err != nil {
		_, err = fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
		if err != nil {
			panic(err)
		}
		os.Exit(1)
	}

	fmt.Println("Running tests....")
	result := m.Run()

	fmt.Println("Cleaning up...")
	err := os.Remove(binName)

	if err != nil {
		panic(err)
	}

	os.Exit(result)
}
