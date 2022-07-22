package temphttp

import (
	"fmt"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/error_handler"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/files"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const envAuthKeysName = "AUTH_KEYS_PATHNAME"

type FiboClient struct {
}

func NewFiboClient() *FiboClient {
	return &FiboClient{}
}

func (f FiboClient) Call(arg string) {
	var (
		serverPort       = os.Getenv("SERVER_PORT")
		apiEndpoint      = os.Getenv("API_ENDPOINT")
		authKeysPathname = os.Getenv(envAuthKeysName)
		requestURL       = fmt.Sprintf("%s:%s/fibonacci/%s", apiEndpoint, serverPort, arg)
	)

	req, requestErr := http.NewRequest(http.MethodGet, requestURL, nil)
	error_handler.AnnotatingError(requestErr, "Failed to create request")

	authKeys := files.ReadFile(authKeysPathname)
	authKey := strings.Split(authKeys, "\n")[0]

	req.Header.Set("Authorization", authKey)

	res, clientErr := http.DefaultClient.Do(req)
	error_handler.AnnotatingError(clientErr, "Failed to send request")

	resBody, readErr := ioutil.ReadAll(res.Body)
	error_handler.AnnotatingError(readErr, "Failed to read response")

	fmt.Printf("fibo %s: %s \n", arg, resBody)
}
