// Code generated by MockGen. DO NOT EDIT.
// Source: ./src/api/services/transactions_service.go

// Package services is a generated GoMock package.
package services

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	dtos "github.com/grosinov/transactions-api/src/api/dtos"
	models "github.com/grosinov/transactions-api/src/api/models"
)

// MockTransactionsService is a mock of TransactionsService interface.
type MockTransactionsService struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionsServiceMockRecorder
}

// MockTransactionsServiceMockRecorder is the mock recorder for MockTransactionsService.
type MockTransactionsServiceMockRecorder struct {
	mock *MockTransactionsService
}

// NewMockTransactionsService creates a new mock instance.
func NewMockTransactionsService(ctrl *gomock.Controller) *MockTransactionsService {
	mock := &MockTransactionsService{ctrl: ctrl}
	mock.recorder = &MockTransactionsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionsService) EXPECT() *MockTransactionsServiceMockRecorder {
	return m.recorder
}

// BulkCreateTransactions mocks base method.
func (m *MockTransactionsService) BulkCreateTransactions(arg0 *[]models.Transaction) (*[]models.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkCreateTransactions", arg0)
	ret0, _ := ret[0].(*[]models.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BulkCreateTransactions indicates an expected call of BulkCreateTransactions.
func (mr *MockTransactionsServiceMockRecorder) BulkCreateTransactions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkCreateTransactions", reflect.TypeOf((*MockTransactionsService)(nil).BulkCreateTransactions), arg0)
}

// GetBalance mocks base method.
func (m *MockTransactionsService) GetBalance(userId uint64, from, to *time.Time) (*dtos.Balance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", userId, from, to)
	ret0, _ := ret[0].(*dtos.Balance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockTransactionsServiceMockRecorder) GetBalance(userId, from, to interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockTransactionsService)(nil).GetBalance), userId, from, to)
}

// GetTransactions mocks base method.
func (m *MockTransactionsService) GetTransactions(userId uint64, from, to *time.Time) (*[]models.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactions", userId, from, to)
	ret0, _ := ret[0].(*[]models.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactions indicates an expected call of GetTransactions.
func (mr *MockTransactionsServiceMockRecorder) GetTransactions(userId, from, to interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactions", reflect.TypeOf((*MockTransactionsService)(nil).GetTransactions), userId, from, to)
}
