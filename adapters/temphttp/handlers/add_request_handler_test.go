package handlers_test

import (
	"bytes"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/temphttp/handlers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestAddRequestHandler(t *testing.T) {
	addRequestHandlerForFormUrlEncodedTests := []struct {
		Name         string
		Body         url.Values
		ResponseBody string
		ResponseCode int
	}{
		{
			Name: "Given one number in Body should return the message with the same number",
			Body: url.Values{
				"num": []string{"2"},
			},
			ResponseBody: "Sum of 2 equal 2 \n",
			ResponseCode: http.StatusOK,
		},
		{
			Name: "Given the wrong Body key should ignore it and give the sum of correct one",
			Body: url.Values{
				"num":      []string{"2", "3"},
				"wrongNum": []string{"20"},
			},
			ResponseBody: "Sum of 2,3 equal 5 \n",
			ResponseCode: http.StatusOK,
		},
		{
			Name: "Given and empty Body should return 400",
			Body: url.Values{
				"num": []string{},
			},
			ResponseBody: "",
			ResponseCode: http.StatusBadRequest,
		},
	}

	for _, tt := range addRequestHandlerForFormUrlEncodedTests {
		t.Run(tt.Name, func(t *testing.T) {
			data := tt.Body
			bodyReader := strings.NewReader(data.Encode())

			request, _ := http.NewRequest(http.MethodPost, "/add", bodyReader)
			response := httptest.NewRecorder()

			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			handlers.AddRequestHandler(response, request)

			gotBody := response.Body.String()
			gotCode := response.Code

			assert.Equal(t, tt.ResponseBody, gotBody)
			assert.Equal(t, tt.ResponseCode, gotCode)
		})
	}

	addRequestHandlerForJsonTests := []struct {
		Name         string
		Body         []byte
		ResponseBody string
		ResponseCode int
	}{
		{
			Name:         "Given one number in Body should return the message with the same number",
			Body:         []byte(`{"nums": [2]}`),
			ResponseBody: "Sum of 2 equal 2 \n",
			ResponseCode: http.StatusOK,
		},
		{
			Name:         "Given the wrong Body key should ignore it and give the sum of correct one",
			Body:         []byte(`{"nums": [2, 3], "wrongNums": 20}`),
			ResponseBody: "Sum of 2,3 equal 5 \n",
			ResponseCode: http.StatusOK,
		},
		{
			Name:         "Given and empty Body should return 400",
			Body:         []byte(`{"nums": []}`),
			ResponseBody: "",
			ResponseCode: 400,
		},
	}

	for _, tt := range addRequestHandlerForJsonTests {
		t.Run(tt.Name, func(t *testing.T) {
			jsonBody := tt.Body
			bodyReader := bytes.NewReader(jsonBody)

			request, _ := http.NewRequest(http.MethodPost, "/add", bodyReader)
			response := httptest.NewRecorder()

			request.Header.Set("Content-Type", "application/json")

			handlers.AddRequestHandler(response, request)

			gotBody := response.Body.String()
			gotCode := response.Code

			assert.Equal(t, tt.ResponseBody, gotBody)
			assert.Equal(t, tt.ResponseCode, gotCode)
		})
	}

	addRequestsHandlersForQueriesTests := []struct {
		Name         string
		queries      string
		ResponseBody string
		ResponseCode int
	}{
		{
			Name:         "Given one number in query should return the message with the same number",
			queries:      "?num=2",
			ResponseBody: "Sum of 2 equal 2 \n",
			ResponseCode: http.StatusOK,
		},
		{
			Name:         "Given the wrong query key only should return 400",
			queries:      "?wrongNum=20",
			ResponseBody: "",
			ResponseCode: http.StatusBadRequest,
		},
		{
			Name:         "Given and empty query should return 400",
			queries:      "",
			ResponseBody: "",
			ResponseCode: http.StatusBadRequest,
		},
	}

	for _, tt := range addRequestsHandlersForQueriesTests {
		t.Run(tt.Name, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodPost, "/add/"+tt.queries, nil)
			response := httptest.NewRecorder()

			handlers.AddRequestHandler(response, request)

			gotBody := response.Body.String()
			gotCode := response.Code

			assert.Equal(t, tt.ResponseBody, gotBody)
			assert.Equal(t, tt.ResponseCode, gotCode)
		})
	}
}
