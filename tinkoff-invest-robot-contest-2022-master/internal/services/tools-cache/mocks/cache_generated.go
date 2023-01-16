// Code generated by MockGen. DO NOT EDIT.
// Source: cache.go

// Package toolscachemocks is a generated GoMock package.
package toolscachemocks

import (
	context "context"
	reflect "reflect"

	tinkoffinvest "github.com/Antonboom/tinkoff-invest-robot-contest-2022/internal/clients/tinkoffinvest"
	gomock "github.com/golang/mock/gomock"
)

// MockSharesProvider is a mock of SharesProvider interface.
type MockSharesProvider struct {
	ctrl     *gomock.Controller
	recorder *MockSharesProviderMockRecorder
}

// MockSharesProviderMockRecorder is the mock recorder for MockSharesProvider.
type MockSharesProviderMockRecorder struct {
	mock *MockSharesProvider
}

// NewMockSharesProvider creates a new mock instance.
func NewMockSharesProvider(ctrl *gomock.Controller) *MockSharesProvider {
	mock := &MockSharesProvider{ctrl: ctrl}
	mock.recorder = &MockSharesProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSharesProvider) EXPECT() *MockSharesProviderMockRecorder {
	return m.recorder
}

// GetShareByFIGI mocks base method.
func (m *MockSharesProvider) GetShareByFIGI(ctx context.Context, figi tinkoffinvest.FIGI) (*tinkoffinvest.Instrument, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShareByFIGI", ctx, figi)
	ret0, _ := ret[0].(*tinkoffinvest.Instrument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShareByFIGI indicates an expected call of GetShareByFIGI.
func (mr *MockSharesProviderMockRecorder) GetShareByFIGI(ctx, figi interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShareByFIGI", reflect.TypeOf((*MockSharesProvider)(nil).GetShareByFIGI), ctx, figi)
}
