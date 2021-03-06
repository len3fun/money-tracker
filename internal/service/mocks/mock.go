// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/len3fun/money-tracker/internal/models"
)

// MockAuthorization is a mock of Authorization interface.
type MockAuthorization struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationMockRecorder
}

// MockAuthorizationMockRecorder is the mock recorder for MockAuthorization.
type MockAuthorizationMockRecorder struct {
	mock *MockAuthorization
}

// NewMockAuthorization creates a new mock instance.
func NewMockAuthorization(ctrl *gomock.Controller) *MockAuthorization {
	mock := &MockAuthorization{ctrl: ctrl}
	mock.recorder = &MockAuthorizationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorization) EXPECT() *MockAuthorizationMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockAuthorization) CreateUser(user models.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAuthorizationMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAuthorization)(nil).CreateUser), user)
}

// GenerateToken mocks base method.
func (m *MockAuthorization) GenerateToken(username, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", username, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockAuthorizationMockRecorder) GenerateToken(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockAuthorization)(nil).GenerateToken), username, password)
}

// ParseToken mocks base method.
func (m *MockAuthorization) ParseToken(token string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", token)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockAuthorizationMockRecorder) ParseToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockAuthorization)(nil).ParseToken), token)
}

// MockMoneySource is a mock of MoneySource interface.
type MockMoneySource struct {
	ctrl     *gomock.Controller
	recorder *MockMoneySourceMockRecorder
}

// MockMoneySourceMockRecorder is the mock recorder for MockMoneySource.
type MockMoneySourceMockRecorder struct {
	mock *MockMoneySource
}

// NewMockMoneySource creates a new mock instance.
func NewMockMoneySource(ctrl *gomock.Controller) *MockMoneySource {
	mock := &MockMoneySource{ctrl: ctrl}
	mock.recorder = &MockMoneySourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMoneySource) EXPECT() *MockMoneySourceMockRecorder {
	return m.recorder
}

// CreateSource mocks base method.
func (m *MockMoneySource) CreateSource(source models.Source) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSource", source)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSource indicates an expected call of CreateSource.
func (mr *MockMoneySourceMockRecorder) CreateSource(source interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSource", reflect.TypeOf((*MockMoneySource)(nil).CreateSource), source)
}

// GetAllSources mocks base method.
func (m *MockMoneySource) GetAllSources(userId int) ([]models.SourceOut, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSources", userId)
	ret0, _ := ret[0].([]models.SourceOut)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSources indicates an expected call of GetAllSources.
func (mr *MockMoneySourceMockRecorder) GetAllSources(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSources", reflect.TypeOf((*MockMoneySource)(nil).GetAllSources), userId)
}

// MockCurrency is a mock of Currency interface.
type MockCurrency struct {
	ctrl     *gomock.Controller
	recorder *MockCurrencyMockRecorder
}

// MockCurrencyMockRecorder is the mock recorder for MockCurrency.
type MockCurrencyMockRecorder struct {
	mock *MockCurrency
}

// NewMockCurrency creates a new mock instance.
func NewMockCurrency(ctrl *gomock.Controller) *MockCurrency {
	mock := &MockCurrency{ctrl: ctrl}
	mock.recorder = &MockCurrencyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCurrency) EXPECT() *MockCurrencyMockRecorder {
	return m.recorder
}

// CreateCurrency mocks base method.
func (m *MockCurrency) CreateCurrency(item models.Currency) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCurrency", item)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCurrency indicates an expected call of CreateCurrency.
func (mr *MockCurrencyMockRecorder) CreateCurrency(item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCurrency", reflect.TypeOf((*MockCurrency)(nil).CreateCurrency), item)
}

// GetAllCurrencies mocks base method.
func (m *MockCurrency) GetAllCurrencies() ([]models.Currency, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCurrencies")
	ret0, _ := ret[0].([]models.Currency)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCurrencies indicates an expected call of GetAllCurrencies.
func (mr *MockCurrencyMockRecorder) GetAllCurrencies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCurrencies", reflect.TypeOf((*MockCurrency)(nil).GetAllCurrencies))
}

// MockActivity is a mock of Activity interface.
type MockActivity struct {
	ctrl     *gomock.Controller
	recorder *MockActivityMockRecorder
}

// MockActivityMockRecorder is the mock recorder for MockActivity.
type MockActivityMockRecorder struct {
	mock *MockActivity
}

// NewMockActivity creates a new mock instance.
func NewMockActivity(ctrl *gomock.Controller) *MockActivity {
	mock := &MockActivity{ctrl: ctrl}
	mock.recorder = &MockActivityMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockActivity) EXPECT() *MockActivityMockRecorder {
	return m.recorder
}

// CreateActivity mocks base method.
func (m *MockActivity) CreateActivity(activity models.Activity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateActivity", activity)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateActivity indicates an expected call of CreateActivity.
func (mr *MockActivityMockRecorder) CreateActivity(activity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateActivity", reflect.TypeOf((*MockActivity)(nil).CreateActivity), activity)
}

// GetAllActivities mocks base method.
func (m *MockActivity) GetAllActivities(userId int) ([]models.ActivitiesOut, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllActivities", userId)
	ret0, _ := ret[0].([]models.ActivitiesOut)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllActivities indicates an expected call of GetAllActivities.
func (mr *MockActivityMockRecorder) GetAllActivities(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllActivities", reflect.TypeOf((*MockActivity)(nil).GetAllActivities), userId)
}
