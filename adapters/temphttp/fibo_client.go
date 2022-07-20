package temphttp

import (
	"fmt"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/error_handler"
	"io/ioutil"
	"net/http"
	"os"
)

type FiboClient struct {
}

func NewFiboClient() *FiboClient {
	return &FiboClient{}
}

func (f FiboClient) Call(arg string) {
	var (
		serverPort  = os.Getenv("SERVER_PORT")
		apiEndpoint = os.Getenv("API_ENDPOINT")
		requestURL  = fmt.Sprintf("%s:%s/fibonacci/%s", apiEndpoint, serverPort, arg)
	)

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	error_handler.HandlePanic(err)

	req.Header.Set("Authorization", "Bearer SUPER_SECRET_API_KEY_1")

	res, err := http.DefaultClient.Do(req)
	error_handler.HandlePanic(err)

	resBody, err := ioutil.ReadAll(res.Body)
	error_handler.HandlePanic(err)

	fmt.Printf("%s", resBody)
}
