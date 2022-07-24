package temphttp

import (
	"fmt"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/files"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const envAuthKeysName = "AUTH_KEYS_PATHNAME"
const envBaseURLName = "API_ENDPOINT"
const envServerPotName = "SERVER_PORT"
const failedToCallFiboClientErrorMessage = "failed to call new request in fibonacci client"

type FiboClient struct {
	baseURL string
	client  *http.Client
}

func NewFiboClient(baseURL string, client *http.Client) *FiboClient {
	return &FiboClient{
		baseURL: baseURL,
		client:  client,
	}
}

func (f FiboClient) Call(arg string) (*http.Request, error) {
	var (
		serverPort       = os.Getenv(envServerPotName)
		apiEndpoint      = os.Getenv(envBaseURLName)
		authKeysPathname = os.Getenv(envAuthKeysName)
		requestURL       = fmt.Sprintf("%s:%s/fibonacci/%s", apiEndpoint, serverPort, arg)
	)

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, errors.Wrap(err, failedToCallFiboClientErrorMessage)
	}

	authKeys, err := files.ReadFile(authKeysPathname)
	if err != nil {
		return nil, errors.Wrap(err, failedToCallFiboClientErrorMessage)
	}

	authKey := strings.Split(authKeys, "\n")[0]
	req.Header.Set("Authorization", authKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, failedToCallFiboClientErrorMessage)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, failedToCallFiboClientErrorMessage)
	}

	fmt.Printf("fibo %s: %s \n", arg, resBody)
	return req, nil
}
