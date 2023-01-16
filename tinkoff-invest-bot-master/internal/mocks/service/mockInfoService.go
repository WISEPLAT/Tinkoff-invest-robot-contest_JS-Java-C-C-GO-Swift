// Code generated by MockGen. DO NOT EDIT.
// Source: ./infoService.go

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	dtotapi "github.com/ldmi3i/tinkoff-invest-bot/internal/dto/dtotapi"
	entity "github.com/ldmi3i/tinkoff-invest-bot/internal/entity"
	investapi "github.com/ldmi3i/tinkoff-invest-bot/internal/tapigen"
)

// MockInfoSrv is a mock of InfoSrv interface.
type MockInfoSrv struct {
	ctrl     *gomock.Controller
	recorder *MockInfoSrvMockRecorder
}

// MockInfoSrvMockRecorder is the mock recorder for MockInfoSrv.
type MockInfoSrvMockRecorder struct {
	mock *MockInfoSrv
}

// NewMockInfoSrv creates a new mock instance.
func NewMockInfoSrv(ctrl *gomock.Controller) *MockInfoSrv {
	mock := &MockInfoSrv{ctrl: ctrl}
	mock.recorder = &MockInfoSrvMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInfoSrv) EXPECT() *MockInfoSrvMockRecorder {
	return m.recorder
}

// GetAccounts mocks base method.
func (m *MockInfoSrv) GetAccounts(ctx context.Context) (*dtotapi.AccountsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccounts", ctx)
	ret0, _ := ret[0].(*dtotapi.AccountsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccounts indicates an expected call of GetAccounts.
func (mr *MockInfoSrvMockRecorder) GetAccounts(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccounts", reflect.TypeOf((*MockInfoSrv)(nil).GetAccounts), ctx)
}

// GetAllShares mocks base method.
func (m *MockInfoSrv) GetAllShares(ctx context.Context) (*dtotapi.SharesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllShares", ctx)
	ret0, _ := ret[0].(*dtotapi.SharesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllShares indicates an expected call of GetAllShares.
func (mr *MockInfoSrvMockRecorder) GetAllShares(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllShares", reflect.TypeOf((*MockInfoSrv)(nil).GetAllShares), ctx)
}

// GetDataStream mocks base method.
func (m *MockInfoSrv) GetDataStream(ctx context.Context) (investapi.MarketDataStreamService_MarketDataStreamClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDataStream", ctx)
	ret0, _ := ret[0].(investapi.MarketDataStreamService_MarketDataStreamClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDataStream indicates an expected call of GetDataStream.
func (mr *MockInfoSrvMockRecorder) GetDataStream(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDataStream", reflect.TypeOf((*MockInfoSrv)(nil).GetDataStream), ctx)
}

// GetHistorySorted mocks base method.
func (m *MockInfoSrv) GetHistorySorted(finis []string, ivl investapi.CandleInterval, startTime, endTime time.Time, ctx context.Context) ([]entity.History, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHistorySorted", finis, ivl, startTime, endTime, ctx)
	ret0, _ := ret[0].([]entity.History)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHistorySorted indicates an expected call of GetHistorySorted.
func (mr *MockInfoSrvMockRecorder) GetHistorySorted(finis, ivl, startTime, endTime, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHistorySorted", reflect.TypeOf((*MockInfoSrv)(nil).GetHistorySorted), finis, ivl, startTime, endTime, ctx)
}

// GetInstrumentInfoByFigi mocks base method.
func (m *MockInfoSrv) GetInstrumentInfoByFigi(figi string, ctx context.Context) (*dtotapi.InstrumentResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstrumentInfoByFigi", figi, ctx)
	ret0, _ := ret[0].(*dtotapi.InstrumentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstrumentInfoByFigi indicates an expected call of GetInstrumentInfoByFigi.
func (mr *MockInfoSrvMockRecorder) GetInstrumentInfoByFigi(figi, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstrumentInfoByFigi", reflect.TypeOf((*MockInfoSrv)(nil).GetInstrumentInfoByFigi), figi, ctx)
}

// GetLastPrices mocks base method.
func (m *MockInfoSrv) GetLastPrices(figis []string, ctx context.Context) (*dtotapi.LastPricesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastPrices", figis, ctx)
	ret0, _ := ret[0].(*dtotapi.LastPricesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLastPrices indicates an expected call of GetLastPrices.
func (mr *MockInfoSrvMockRecorder) GetLastPrices(figis, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastPrices", reflect.TypeOf((*MockInfoSrv)(nil).GetLastPrices), figis, ctx)
}

// GetOrderState mocks base method.
func (m *MockInfoSrv) GetOrderState(req *dtotapi.OrderStateRequest, ctx context.Context) (*dtotapi.OrderStateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderState", req, ctx)
	ret0, _ := ret[0].(*dtotapi.OrderStateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderState indicates an expected call of GetOrderState.
func (mr *MockInfoSrvMockRecorder) GetOrderState(req, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderState", reflect.TypeOf((*MockInfoSrv)(nil).GetOrderState), req, ctx)
}

// GetPositions mocks base method.
func (m *MockInfoSrv) GetPositions(req *dtotapi.PositionsRequest, ctx context.Context) (*dtotapi.PositionsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPositions", req, ctx)
	ret0, _ := ret[0].(*dtotapi.PositionsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPositions indicates an expected call of GetPositions.
func (mr *MockInfoSrvMockRecorder) GetPositions(req, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPositions", reflect.TypeOf((*MockInfoSrv)(nil).GetPositions), req, ctx)
}