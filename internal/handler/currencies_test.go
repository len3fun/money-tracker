package handler

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/len3fun/money-tracker/internal/models"
	"github.com/len3fun/money-tracker/internal/service"
	mock_service "github.com/len3fun/money-tracker/internal/service/mocks"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_createCurrency(t *testing.T) {
	type mockBehavior func(s *mock_service.MockCurrency, inputCurrency models.Currency)

	testTable := []struct {
		name                string
		inputBody           string
		inputCurrency       models.Currency
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name": "Test", "ticket": "Test"}`,
			inputCurrency: models.Currency{
				Name:   "Test",
				Ticket: "Test",
			},
			mockBehavior: func(s *mock_service.MockCurrency, inputCurrency models.Currency) {
				s.EXPECT().CreateCurrency(inputCurrency).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name:                "Empty fields",
			inputBody:           `{"ticket": "Test"}`,
			mockBehavior:        func(s *mock_service.MockCurrency, inputCurrency models.Currency) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"invalid currency input, fields 'name' and 'ticket' are required"}`,
		},
		{
			name:      "Service failure",
			inputBody: `{"name": "Test", "ticket": "Test"}`,
			inputCurrency: models.Currency{
				Name:   "Test",
				Ticket: "Test",
			},
			mockBehavior: func(s *mock_service.MockCurrency, inputCurrency models.Currency) {
				s.EXPECT().CreateCurrency(inputCurrency).Return(0, errors.New("service failure"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"service failure"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			currency := mock_service.NewMockCurrency(c)
			testCase.mockBehavior(currency, testCase.inputCurrency)

			services := &service.Service{Currency: currency}
			handler := NewHandler(services)

			r := gin.New()
			r.POST("/currencies", handler.createCurrency)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/currencies",
				bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})

	}
}

func TestHandler_getAllCurrencies(t *testing.T) {
	type mockBehavior func(s *mock_service.MockCurrency)

	testTable := []struct {
		name                string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name: "OK",
			mockBehavior: func(s *mock_service.MockCurrency) {
				s.EXPECT().GetAllCurrencies().Return(
					[]models.Currency{{1, "test1", "t1"}, {2, "test2", "t2"}}, nil,
				)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `[{"id":1,"name":"test1","ticket":"t1"},{"id":2,"name":"test2","ticket":"t2"}]`,
		},
		{
			name: "Service failure",
			mockBehavior: func(s *mock_service.MockCurrency) {
				s.EXPECT().GetAllCurrencies().Return(
					[]models.Currency{}, errors.New("service failure"),
				)
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"service failure"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			currency := mock_service.NewMockCurrency(c)
			testCase.mockBehavior(currency)

			services := &service.Service{Currency: currency}
			handler := NewHandler(services)

			r := gin.New()
			r.GET("/currencies", handler.getAllCurrencies)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/currencies", bytes.NewBufferString(""))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
