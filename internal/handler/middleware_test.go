package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/len3fun/money-tracker/internal/service"
	mock_service "github.com/len3fun/money-tracker/internal/service/mocks"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_userIdentity(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorization, token string)

	testTable := []struct {
		name                 string
		headerName           string
		headerValue          string
		token                string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	} {
		{
			name: "OK",
			headerName: "Authorization",
			headerValue: "Bearer token",
			token: "token",
			mockBehavior: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, nil)
			},
			expectedStatusCode: 200,
			expectedResponseBody: "1",
		},
		{
			name: "No header",
			headerName: "",
			mockBehavior: func(s *mock_service.MockAuthorization, token string) {},
			expectedStatusCode: 401,
			expectedResponseBody: `{"message":"empty auth header"}`,
		},
		{
			name: "Invalid Bearer",
			headerName: "Authorization",
			headerValue: "Bearr token",
			mockBehavior: func(s *mock_service.MockAuthorization, token string) {},
			expectedStatusCode: 401,
			expectedResponseBody: `{"message":"invalid auth header"}`,
		},
		{
			name: "Invalid token",
			headerName: "Authorization",
			headerValue: "Bearer ",
			mockBehavior: func(s *mock_service.MockAuthorization, token string) {},
			expectedStatusCode: 401,
			expectedResponseBody: `{"message":"token is empty"}`,
		},
		{
			name: "Service Failure",
			headerName: "Authorization",
			headerValue: "Bearer token",
			token: "token",
			mockBehavior: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, errors.New("failed to parse token"))
			},
			expectedStatusCode: 401,
			expectedResponseBody: `{"message":"failed to parse token"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehavior(auth, testCase.token)

			services := &service.Service{Authorization: auth}
			handler := NewHandler(services)

			r := gin.New()
			r.GET("/protected", handler.userIdentity, func(c *gin.Context) {
				id, _ := c.Get(userCtx)
				c.String(200, fmt.Sprintf("%d", id.(int)))
			})

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/protected", nil)
			req.Header.Set(testCase.headerName, testCase.headerValue)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedResponseBody)
		})
	}
}
