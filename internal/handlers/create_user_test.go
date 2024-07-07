package handlers

import (
	"effectiveMobileTestProblem/internal/handlers/mocks"
	"effectiveMobileTestProblem/internal/model"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateUser(t *testing.T) {
	type mockBehavior func(service *mocks.UserService)

	var tests = []struct {
		name         string
		body         string
		mockBehavior mockBehavior
		wantAnswer   string
		wantStatus   int
	}{
		{
			name: "Good test",
			body: `{
	"passport_series_and_number": "1234 567894",
    "name": "Zhan211111",
    "surname": "Zhakushev", 
    "address": "Moscow"
}`,
			mockBehavior: func(service *mocks.UserService) {
				service.On("CreateUser", mock.Anything, mock.Anything).Return("id", nil)
			},
			wantAnswer: `{"id":"id"}`,
			wantStatus: http.StatusCreated,
		}, {
			name: "Incorrect request body",
			body: `{
	"passport_series_and_number": "1234 567894",
	"name": "Zhan211111",
	"surname": "Zhakushev"
	"address": "Moscow"
	}`,
			mockBehavior: func(service *mocks.UserService) {},
			wantAnswer:   `{"message":"Invalid request. Check request body"}`,
			wantStatus:   http.StatusBadRequest,
		}, {
			name: "Incorrect len of passportData",
			body: `{
	"passport_series_and_number": "1234",
	"name": "Zhan211111",
	"surname": "Zhakushev",
	"address": "Moscow"
}`,
			mockBehavior: func(service *mocks.UserService) {},
			wantAnswer:   `{"message":"Invalid passport series and number format. Should be in next format: 1234 123456"}`,
			wantStatus:   http.StatusBadRequest,
		}, {
			name: "Incorrect passport number",
			body: `{
	"passport_series_and_number": "1234 a123456",
	"name": "Zhan211111",
	"surname": "Zhakushev",
	"address": "Moscow"
}`,
			mockBehavior: func(service *mocks.UserService) {},
			wantAnswer:   `{"message":"Invalid passport number. Passport number should be a number"}`,
			wantStatus:   http.StatusBadRequest,
		}, {
			name: "Incorrect passport series",
			body: `{
	"passport_series_and_number": "1234a 123456",
	"name": "Zhan211111",
	"surname": "Zhakushev",
	"address": "Moscow"
}`,
			mockBehavior: func(service *mocks.UserService) {},
			wantAnswer:   `{"message":"Invalid passport series. Passport series should be a number"}`,
			wantStatus:   http.StatusBadRequest,
		}, {
			name: "User already exists",
			body: `{
	"passport_series_and_number": "1234 567894",
	"name": "Zhan211111",
	"surname": "Zhakushev",
	"address": "Moscow"
}`,
			mockBehavior: func(service *mocks.UserService) {
				service.On("CreateUser", mock.Anything, mock.Anything).Return("", model.ErrAlreadyExists)
			},
			wantAnswer: `{"message":"User with this passport series and number already exists"}`,
			wantStatus: http.StatusBadRequest,
		}, {
			name: "Create user error",
			body: `{
	"passport_series_and_number": "1234 567894",
	"name": "Zhan211111",
	"surname": "Zhakushev",
	"address": "Moscow"
}`,
			mockBehavior: func(service *mocks.UserService) {
				service.On("CreateUser", mock.Anything, mock.Anything).Return("", assert.AnError)
			},
			wantAnswer: `{"message":"Failed to create user"}`,
			wantStatus: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			serviceMock := mocks.NewUserService(t)

			test.mockBehavior(serviceMock)

			e := echo.New()

			NewHandlers(e, serviceMock, nil)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/api/user", strings.NewReader(test.body))
			r.Header.Set("Content-Type", "application/json")

			e.ServeHTTP(w, r)

			assert.Equal(t, test.wantStatus, w.Code)
			assert.JSONEq(t, test.wantAnswer, w.Body.String())

		})

	}
}
