package handlers_test

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/temphttp/handlers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestAddRequestHandlerForFormUrlEncoded(t *testing.T) {
	addRequestHandlerForFormUrlEncodedTests := []struct {
		Name         string
		body         url.Values
		responseBody string
		responseCode int
	}{
		{
			Name: "Given one number in body should return the message with the same number",
			body: url.Values{
				"num": []string{"2"},
			},
			responseBody: "Sum of 2 equal 2 \n",
			responseCode: http.StatusOK,
		},
		{
			Name: "Given two numbers in body should return the message with the correct sum of them",
			body: url.Values{
				"num": []string{"2", "3"},
			},
			responseBody: "Sum of 2,3 equal 5 \n",
			responseCode: http.StatusOK,
		},
		{
			Name: "Given the wrong body key should ignore it and give the sum of correct one",
			body: url.Values{
				"num":      []string{"2", "3"},
				"wrongNum": []string{"20"},
			},
			responseBody: "Sum of 2,3 equal 5 \n",
			responseCode: http.StatusOK,
		},
		{
			Name: "Given the wrong body key only should return 400",
			body: url.Values{
				"wrongNum": []string{"2", "3"},
			},
			responseBody: "",
			responseCode: http.StatusBadRequest,
		},
		{
			Name: "Given and empty body should return 400",
			body: url.Values{
				"num": []string{},
			},
			responseBody: "",
			responseCode: http.StatusBadRequest,
		},
	}

	for _, tt := range addRequestHandlerForFormUrlEncodedTests {
		t.Run(tt.Name, func(t *testing.T) {
			data := tt.body
			bodyReader := strings.NewReader(data.Encode())

			request, _ := http.NewRequest(http.MethodPost, "/add", bodyReader)
			response := httptest.NewRecorder()

			request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			handlers.AddRequestHandlerForFormUrlEncoded(response, request)

			gotBody := response.Body.String()
			gotCode := response.Code

			assert.Equal(t, tt.responseBody, gotBody)
			assert.Equal(t, tt.responseCode, gotCode)
		})
	}
}
