package handlers_test

import (
	"bytes"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/temphttp/handlers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddRequestHandlerForJson(t *testing.T) {
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
			Name:         "Given two numbers in Body should return the message with the correct sum of them",
			Body:         []byte(`{"nums": [2, 3]}`),
			ResponseBody: "Sum of 2,3 equal 5 \n",
			ResponseCode: http.StatusOK,
		},
		{
			Name:         "Given the wrong Body key should ignore it and give the sum of correct one",
			Body:         []byte(`{"nums": [2, 3], "wrongNums": 20}`),
			ResponseBody: "Sum of 2,3 equal 5 \n",
			ResponseCode: http.StatusOK,
		},
		{
			Name:         "Given the wrong Body key only should return 400",
			Body:         []byte(`{"wrongNums": ["2", "3"]}`),
			ResponseBody: "",
			ResponseCode: 400,
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

			handlers.AddRequestHandlerForJson(response, request)

			gotBody := response.Body.String()
			gotCode := response.Code

			assert.Equal(t, tt.ResponseBody, gotBody)
			assert.Equal(t, tt.ResponseCode, gotCode)
		})
	}
}
