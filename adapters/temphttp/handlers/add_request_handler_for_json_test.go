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
		body         []byte
		responseBody string
		responseCode int
	}{
		{
			Name:         "Given one number in body should return the message with the same number",
			body:         []byte(`{"nums": [2]}`),
			responseBody: "Sum of 2 equal 2 \n",
			responseCode: 200,
		},
		{
			Name:         "Given two numbers in body should return the message with the correct sum of them",
			body:         []byte(`{"nums": [2, 3]}`),
			responseBody: "Sum of 2,3 equal 5 \n",
			responseCode: 200,
		},
		{
			Name:         "Given the wrong body key should ignore it and give the sum of correct one",
			body:         []byte(`{"nums": [2, 3], "wrongNums": 20}`),
			responseBody: "Sum of 2,3 equal 5 \n",
			responseCode: 200,
		},
		{
			Name:         "Given the wrong body key only should return 400",
			body:         []byte(`{"wrongNums": ["2", "3"]}`),
			responseBody: "",
			responseCode: 400,
		},
		{
			Name:         "Given and empty body should return 400",
			body:         []byte(`{"nums": []}`),
			responseBody: "",
			responseCode: 400,
		},
	}

	for _, tt := range addRequestHandlerForJsonTests {
		t.Run(tt.Name, func(t *testing.T) {
			jsonBody := tt.body
			bodyReader := bytes.NewReader(jsonBody)

			request, _ := http.NewRequest(http.MethodPost, "/add", bodyReader)
			response := httptest.NewRecorder()

			request.Header.Set("Content-Type", "application/json")

			handlers.AddRequestHandlerForJson(response, request)

			gotBody := response.Body.String()
			gotCode := response.Code

			assert.Equal(t, tt.responseBody, gotBody)
			assert.Equal(t, tt.responseCode, gotCode)
		})
	}
}
