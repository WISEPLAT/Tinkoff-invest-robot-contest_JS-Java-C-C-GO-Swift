// Code generated by MockGen. DO NOT EDIT.
// Source: ./client.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/liderman/traderstack/internal/domain"
	decimal "github.com/shopspring/decimal"
)

// MockApiClient is a mock of ApiClient interface.
type MockApiClient struct {
	ctrl     *gomock.Controller
	recorder *MockApiClientMockRecorder
}

// MockApiClientMockRecorder is the mock recorder for MockApiClient.
type MockApiClientMockRecorder struct {
	mock *MockApiClient
}

// NewMockApiClient creates a new mock instance.
func NewMockApiClient(ctrl *gomock.Controller) *MockApiClient {
	mock := &MockApiClient{ctrl: ctrl}
	mock.recorder = &MockApiClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApiClient) EXPECT() *MockApiClientMockRecorder {
	return m.recorder
}

// GetAccounts mocks base method.
func (m *MockApiClient) GetAccounts() ([]*domain.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccounts")
	ret0, _ := ret[0].([]*domain.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccounts indicates an expected call of GetAccounts.
func (mr *MockApiClientMockRecorder) GetAccounts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccounts", reflect.TypeOf((*MockApiClient)(nil).GetAccounts))
}

// GetCandles mocks base method.
func (m *MockApiClient) GetCandles(figi string, from, to time.Time, interval domain.CandleInterval) ([]*domain.HistoricCandle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCandles", figi, from, to, interval)
	ret0, _ := ret[0].([]*domain.HistoricCandle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCandles indicates an expected call of GetCandles.
func (mr *MockApiClientMockRecorder) GetCandles(figi, from, to, interval interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCandles", reflect.TypeOf((*MockApiClient)(nil).GetCandles), figi, from, to, interval)
}

// GetLastPrices mocks base method.
func (m *MockApiClient) GetLastPrices() ([]*domain.LastPrice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastPrices")
	ret0, _ := ret[0].([]*domain.LastPrice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLastPrices indicates an expected call of GetLastPrices.
func (mr *MockApiClientMockRecorder) GetLastPrices() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastPrices", reflect.TypeOf((*MockApiClient)(nil).GetLastPrices))
}

// GetOrders mocks base method.
func (m *MockApiClient) GetOrders(accountId string) ([]*domain.OrderState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrders", accountId)
	ret0, _ := ret[0].([]*domain.OrderState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrders indicates an expected call of GetOrders.
func (mr *MockApiClientMockRecorder) GetOrders(accountId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrders", reflect.TypeOf((*MockApiClient)(nil).GetOrders), accountId)
}

// GetPortfolio mocks base method.
func (m *MockApiClient) GetPortfolio(accountId string) ([]*domain.PortfolioPosition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPortfolio", accountId)
	ret0, _ := ret[0].([]*domain.PortfolioPosition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPortfolio indicates an expected call of GetPortfolio.
func (mr *MockApiClientMockRecorder) GetPortfolio(accountId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPortfolio", reflect.TypeOf((*MockApiClient)(nil).GetPortfolio), accountId)
}

// GetSandboxAccounts mocks base method.
func (m *MockApiClient) GetSandboxAccounts() ([]*domain.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSandboxAccounts")
	ret0, _ := ret[0].([]*domain.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSandboxAccounts indicates an expected call of GetSandboxAccounts.
func (mr *MockApiClientMockRecorder) GetSandboxAccounts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSandboxAccounts", reflect.TypeOf((*MockApiClient)(nil).GetSandboxAccounts))
}

// GetSandboxOrders mocks base method.
func (m *MockApiClient) GetSandboxOrders(accountId string) ([]*domain.OrderState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSandboxOrders", accountId)
	ret0, _ := ret[0].([]*domain.OrderState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSandboxOrders indicates an expected call of GetSandboxOrders.
func (mr *MockApiClientMockRecorder) GetSandboxOrders(accountId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSandboxOrders", reflect.TypeOf((*MockApiClient)(nil).GetSandboxOrders), accountId)
}

// GetSandboxPortfolio mocks base method.
func (m *MockApiClient) GetSandboxPortfolio(accountId string) ([]*domain.PortfolioPosition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSandboxPortfolio", accountId)
	ret0, _ := ret[0].([]*domain.PortfolioPosition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSandboxPortfolio indicates an expected call of GetSandboxPortfolio.
func (mr *MockApiClientMockRecorder) GetSandboxPortfolio(accountId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSandboxPortfolio", reflect.TypeOf((*MockApiClient)(nil).GetSandboxPortfolio), accountId)
}

// GetSandboxPositions mocks base method.
func (m *MockApiClient) GetSandboxPositions(accountId string) (*domain.PositionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSandboxPositions", accountId)
	ret0, _ := ret[0].(*domain.PositionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSandboxPositions indicates an expected call of GetSandboxPositions.
func (mr *MockApiClientMockRecorder) GetSandboxPositions(accountId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSandboxPositions", reflect.TypeOf((*MockApiClient)(nil).GetSandboxPositions), accountId)
}

// GetShares mocks base method.
func (m *MockApiClient) GetShares() ([]*domain.Share, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShares")
	ret0, _ := ret[0].([]*domain.Share)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShares indicates an expected call of GetShares.
func (mr *MockApiClientMockRecorder) GetShares() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShares", reflect.TypeOf((*MockApiClient)(nil).GetShares))
}

// OpenSandboxAccount mocks base method.
func (m *MockApiClient) OpenSandboxAccount() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenSandboxAccount")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenSandboxAccount indicates an expected call of OpenSandboxAccount.
func (mr *MockApiClientMockRecorder) OpenSandboxAccount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenSandboxAccount", reflect.TypeOf((*MockApiClient)(nil).OpenSandboxAccount))
}

// PostOrder mocks base method.
func (m *MockApiClient) PostOrder(figi string, lots int64, price decimal.Decimal, direction domain.OrderDirection, accountId string, orderType domain.OrderType) (*domain.PostOrderResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostOrder", figi, lots, price, direction, accountId, orderType)
	ret0, _ := ret[0].(*domain.PostOrderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostOrder indicates an expected call of PostOrder.
func (mr *MockApiClientMockRecorder) PostOrder(figi, lots, price, direction, accountId, orderType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostOrder", reflect.TypeOf((*MockApiClient)(nil).PostOrder), figi, lots, price, direction, accountId, orderType)
}

// PostSandboxOrder mocks base method.
func (m *MockApiClient) PostSandboxOrder(figi string, lots int64, price decimal.Decimal, direction domain.OrderDirection, accountId string, orderType domain.OrderType) (*domain.PostOrderResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostSandboxOrder", figi, lots, price, direction, accountId, orderType)
	ret0, _ := ret[0].(*domain.PostOrderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostSandboxOrder indicates an expected call of PostSandboxOrder.
func (mr *MockApiClientMockRecorder) PostSandboxOrder(figi, lots, price, direction, accountId, orderType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostSandboxOrder", reflect.TypeOf((*MockApiClient)(nil).PostSandboxOrder), figi, lots, price, direction, accountId, orderType)
}

// SandboxPayIn mocks base method.
func (m *MockApiClient) SandboxPayIn(accountId string, amount *domain.MoneyValue) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SandboxPayIn", accountId, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// SandboxPayIn indicates an expected call of SandboxPayIn.
func (mr *MockApiClientMockRecorder) SandboxPayIn(accountId, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SandboxPayIn", reflect.TypeOf((*MockApiClient)(nil).SandboxPayIn), accountId, amount)
}
