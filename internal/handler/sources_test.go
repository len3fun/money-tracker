package handler

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/len3fun/money-tracker/internal/models"
	"github.com/len3fun/money-tracker/internal/service"
	mock_service "github.com/len3fun/money-tracker/internal/service/mocks"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_createSource(t *testing.T) {
	type mockBehavior func(s *mock_service.MockMoneySource, inputSource models.Source)
	type mockAuth func(s *mock_service.MockAuthorization, token string)

	testTable := []struct {
		name                string
		inputBody           string
		inputSource         models.Source
		mockBehavior        mockBehavior
		mockAuth            mockAuth
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"type": "Test", "balance": 1000, "currency_id": 1}`,
			inputSource: models.Source{
				UserId:     1,
				Type:       "Test",
				Balance:    decimal.New(1000, 0),
				CurrencyId: 1,
			},
			mockAuth: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, nil)
			},
			mockBehavior: func(s *mock_service.MockMoneySource, inputSource models.Source) {
				s.EXPECT().CreateSource(inputSource).Return(nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `"Ok"`,
		},
		{
			name:      "Empty type field",
			inputBody: `{"balance": 1000, "currency_id": 1}`,
			mockAuth: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, nil)
			},
			mockBehavior:        func(s *mock_service.MockMoneySource, inputSource models.Source) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"'type' field shouldn't be empty"}`,
		},
		{
			name:      "Empty balance field",
			inputBody: `{"type": "Test", "currency_id": 1}`,
			inputSource: models.Source{
				UserId:     1,
				Type:       "Test",
				CurrencyId: 1,
			},
			mockAuth: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, nil)
			},
			mockBehavior: func(s *mock_service.MockMoneySource, inputSource models.Source) {
				s.EXPECT().CreateSource(inputSource).Return(nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `"Ok"`,
		},
		{
			name:      "Invalid currency_id field",
			inputBody: `{"type": "Test", "balance": 1000}`,
			mockAuth: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, nil)
			},
			mockBehavior:        func(s *mock_service.MockMoneySource, inputSource models.Source) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"'currency_id' field should be greater than 0"}`,
		},
		{
			name:      "Bad user id",
			inputBody: `{"type": "Test", "balance": 1000, "currency_id": 1}`,
			mockAuth: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, errors.New("bad user id"))
			},
			mockBehavior:        func(s *mock_service.MockMoneySource, inputSource models.Source) {},
			expectedStatusCode:  401,
			expectedRequestBody: `{"message":"bad user id"}`,
		},
		{
			name:      "Service failure",
			inputBody: `{"type": "Test", "balance": 1000, "currency_id": 1}`,
			inputSource: models.Source{
				UserId:     1,
				Type:       "Test",
				Balance:    decimal.New(1000, 0),
				CurrencyId: 1,
			},
			mockAuth: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, nil)
			},
			mockBehavior: func(s *mock_service.MockMoneySource, inputSource models.Source) {
				s.EXPECT().CreateSource(inputSource).Return(errors.New("service failure"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"service failure"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			source := mock_service.NewMockMoneySource(c)
			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehavior(source, testCase.inputSource)
			testCase.mockAuth(auth, "token")

			services := &service.Service{MoneySource: source, Authorization: auth}
			handler := NewHandler(services)

			r := gin.New()
			api := r.Group("", handler.userIdentity)
			{
				api.POST("/sources", handler.createSource)
			}

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sources",
				bytes.NewBufferString(testCase.inputBody))
			req.Header.Set(authorizationHeader, "Bearer token")

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})

	}
}

func TestHandler_getAllSources(t *testing.T) {
	type mockBehavior func(s *mock_service.MockMoneySource)
	type mockAuth func(s *mock_service.MockAuthorization, token string)

	testTable := []struct {
		name                string
		mockBehavior        mockBehavior
		mockAuth            mockAuth
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name: "OK",
			mockBehavior: func(s *mock_service.MockMoneySource) {
				s.EXPECT().GetAllSources(1).Return([]models.SourceOut{
					{
						Type:    "test",
						Balance: decimal.New(10, 0),
						Ticket:  "RUB",
					},
					{
						Type:    "Cash",
						Balance: decimal.New(20, 0),
						Ticket:  "RUB",
					},
				}, nil,
				)
			},
			mockAuth: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `[{"type":"test","balance":"10","currency_id":"RUB"},{"type":"Cash","balance":"20","currency_id":"RUB"}]`,
		},
		{
			name:         "Bad user id",
			mockBehavior: func(s *mock_service.MockMoneySource) {},
			mockAuth: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, errors.New("bad user id"))
			},
			expectedStatusCode:  401,
			expectedRequestBody: `{"message":"bad user id"}`,
		},
		{
			name: "Service failure",
			mockBehavior: func(s *mock_service.MockMoneySource) {
				s.EXPECT().GetAllSources(1).Return([]models.SourceOut{},
					errors.New("service failure"),
				)
			},
			mockAuth: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, nil)
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"service failure"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			moneySource := mock_service.NewMockMoneySource(c)
			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehavior(moneySource)
			testCase.mockAuth(auth, "token")

			services := &service.Service{MoneySource: moneySource, Authorization: auth}
			handler := NewHandler(services)

			r := gin.New()
			api := r.Group("", handler.userIdentity)
			{
				api.GET("/sources", handler.getAllSources)
			}

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/sources", bytes.NewBufferString(""))
			req.Header.Set(authorizationHeader, "Bearer token")

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
