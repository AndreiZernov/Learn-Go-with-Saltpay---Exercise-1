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

	req, requestErr := http.NewRequest(http.MethodGet, requestURL, nil)
	error_handler.AnnotatingError(requestErr, "Failed to create request")

	req.Header.Set("Authorization", "Bearer SUPER_SECRET_API_KEY_1")

	res, clientErr := http.DefaultClient.Do(req)
	error_handler.AnnotatingError(clientErr, "Failed to send request")

	resBody, readErr := ioutil.ReadAll(res.Body)
	error_handler.AnnotatingError(readErr, "Failed to read response")

	fmt.Printf("fibo %s: %s \n", arg, resBody)
}
