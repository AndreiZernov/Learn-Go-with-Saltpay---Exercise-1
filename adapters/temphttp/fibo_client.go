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

type FiboClient struct {
}

func NewFiboClient() *FiboClient {
	return &FiboClient{}
}

func (f FiboClient) Call(arg string) error {
	var (
		serverPort       = os.Getenv("SERVER_PORT")
		apiEndpoint      = os.Getenv("API_ENDPOINT")
		authKeysPathname = os.Getenv(envAuthKeysName)
		requestURL       = fmt.Sprintf("%s:%s/fibonacci/%s", apiEndpoint, serverPort, arg)
	)

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return errors.Wrap(err, "failed to call new request in fibonacci client")
	}

	authKeys, err := files.ReadFile(authKeysPathname)
	if err != nil {
		return errors.Wrap(err, "failed to call new request in fibonacci client")
	}

	authKey := strings.Split(authKeys, "\n")[0]

	req.Header.Set("Authorization", authKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to call new request in fibonacci client")
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.Wrap(err, "failed to call new request in fibonacci client")
	}

	fmt.Printf("fibo %s: %s \n", arg, resBody)
	return nil
}
