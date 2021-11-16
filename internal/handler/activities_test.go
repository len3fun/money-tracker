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
	"time"
)

func TestHandler_getAllActivities(t *testing.T) {
	type mockBehavior func(s *mock_service.MockActivity)
	type mockAuth func(s *mock_service.MockAuthorization, token string)

	location, _ := time.LoadLocation("UTC")

	testTable := []struct {
		name                string
		mockBehavior        mockBehavior
		mockAuth            mockAuth
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name: "OK",
			mockAuth: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, nil)
			},
			mockBehavior: func(s *mock_service.MockActivity) {
				s.EXPECT().GetAllActivities(1).Return([]models.ActivitiesOut{
					{
						Type:         "income",
						Change:       decimal.New(1, 0),
						Label:        "test",
						ActivityDate: time.Date(2021, 11, 15, 9, 17, 7, 942137000, location),
					},
					{
						Type:         "income",
						Change:       decimal.New(2, 0),
						Label:        "test2",
						ActivityDate: time.Date(2021, 11, 15, 9, 23, 51, 501018000, location),
					},
				}, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `[{"type":"income","change":"1","label":"test","activity_date":"2021-11-15T09:17:07.942137Z"},{"type":"income","change":"2","label":"test2","activity_date":"2021-11-15T09:23:51.501018Z"}]`,
		},
		{
			name: "Bad user id",
			mockAuth: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(0, errors.New("bad user id"))
			},
			mockBehavior:        func(s *mock_service.MockActivity) {},
			expectedStatusCode:  401,
			expectedRequestBody: `{"message":"bad user id"}`,
		},
		{
			name: "Service failure",
			mockAuth: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, nil)
			},
			mockBehavior: func(s *mock_service.MockActivity) {
				s.EXPECT().GetAllActivities(1).Return([]models.ActivitiesOut{}, errors.New("service failure"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"service failure"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			activities := mock_service.NewMockActivity(c)
			auth := mock_service.NewMockAuthorization(c)

			testCase.mockBehavior(activities)
			testCase.mockAuth(auth, "token")

			services := &service.Service{Activity: activities, Authorization: auth}
			handler := NewHandler(services)

			r := gin.New()
			api := r.Group("", handler.userIdentity)
			{
				api.GET("/activities", handler.getAllActivities)
			}

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/activities", bytes.NewBufferString(""))
			req.Header.Set(authorizationHeader, "Bearer token")

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}

func TestHandler_createActivity(t *testing.T) {
	type mockBehavior func(s *mock_service.MockActivity, inputActivity models.Activity)
	type mockAuth func(s *mock_service.MockAuthorization, token string)

	location, _ := time.LoadLocation("UTC")
	activityDate := time.Date(2021, 11, 16, 11, 47, 20, 942137000, location)

	testTable := []struct {
		name                string
		inputBody           string
		inputActivity       models.Activity
		mockAuth            mockAuth
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"type": "expense", "label": "test", "change": 1, "source_id": 1, "activity_date": "2021-11-16T11:47:20.942137Z"}`,
			inputActivity: models.Activity{
				UserId:       1,
				SourceId:     1,
				Type:         "expense",
				Change:       decimal.New(1, 0),
				Label:        "test",
				ActivityDate: activityDate,
			},
			mockAuth: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, nil)
			},
			mockBehavior: func(s *mock_service.MockActivity, inputActivity models.Activity) {
				s.EXPECT().CreateActivity(inputActivity).Return(nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `"Ok"`,
		},
		{
			name:      "Empty type",
			inputBody: `{"label": "test", "change": 1, "source_id": 1, "activity_date": "2021-11-16T11:47:20.942137Z"}`,
			mockAuth: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, nil)
			},
			mockBehavior: func(s *mock_service.MockActivity, inputActivity models.Activity) {
			},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"field 'type' mustn't be empty"}`,
		},
		{
			name:      "Invalid type",
			inputBody: `{"type": "invalid", "label": "test", "change": 1, "source_id": 1, "activity_date": "2021-11-16T11:47:20.942137Z"}`,
			mockAuth: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, nil)
			},
			mockBehavior: func(s *mock_service.MockActivity, inputActivity models.Activity) {
			},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"field 'type' must be equal to 'income' or 'expense'"}`,
		},
		{
			name:      "Empty label",
			inputBody: `{"type": "income", "change": 1, "source_id": 1, "activity_date": "2021-11-16T11:47:20.942137Z"}`,
			mockAuth: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, nil)
			},
			mockBehavior: func(s *mock_service.MockActivity, inputActivity models.Activity) {
			},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"field 'label' mustn't be empty"}`,
		},
		{
			name:      "Empty change",
			inputBody: `{"type": "income", "label":"test", "source_id": 1, "activity_date": "2021-11-16T11:47:20.942137Z"}`,
			mockAuth: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, nil)
			},
			mockBehavior: func(s *mock_service.MockActivity, inputActivity models.Activity) {
			},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"field 'change' must be greater than zero"}`,
		},
		{
			name:      "Empty source_id",
			inputBody: `{"type": "income", "label":"test", "change": 1, "activity_date": "2021-11-16T11:47:20.942137Z"}`,
			mockAuth: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, nil)
			},
			mockBehavior: func(s *mock_service.MockActivity, inputActivity models.Activity) {
			},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"field 'source_id' mustn't be empty"}`,
		},
		{
			name:      "Bad user id",
			inputBody: `{"type": "income", "label":"test", "change": 1, "source_id", "activity_date": "2021-11-16T11:47:20.942137Z"}`,
			mockAuth: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, errors.New("bad user id"))
			},
			mockBehavior: func(s *mock_service.MockActivity, inputActivity models.Activity) {
			},
			expectedStatusCode:  401,
			expectedRequestBody: `{"message":"bad user id"}`,
		},
		{
			name:      "Service failure",
			inputBody: `{"type": "expense", "label": "test", "change": 1, "source_id": 1, "activity_date": "2021-11-16T11:47:20.942137Z"}`,
			inputActivity: models.Activity{
				UserId:       1,
				SourceId:     1,
				Type:         "expense",
				Change:       decimal.New(1, 0),
				Label:        "test",
				ActivityDate: activityDate,
			},
			mockAuth: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(1, nil)
			},
			mockBehavior: func(s *mock_service.MockActivity, inputActivity models.Activity) {
				s.EXPECT().CreateActivity(inputActivity).Return(errors.New("service failure"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"service failure"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			activities := mock_service.NewMockActivity(c)
			auth := mock_service.NewMockAuthorization(c)

			testCase.mockBehavior(activities, testCase.inputActivity)
			testCase.mockAuth(auth, "token")

			services := &service.Service{Activity: activities, Authorization: auth}
			handler := NewHandler(services)

			r := gin.New()
			api := r.Group("", handler.userIdentity)
			{
				api.POST("/activities", handler.createActivity)
			}

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/activities", bytes.NewBufferString(testCase.inputBody))
			req.Header.Set(authorizationHeader, "Bearer token")

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
