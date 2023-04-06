// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/usecase/interface/auth.interface.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAuthUseCase is a mock of AuthUseCase interface.
type MockAuthUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockAuthUseCaseMockRecorder
}

// MockAuthUseCaseMockRecorder is the mock recorder for MockAuthUseCase.
type MockAuthUseCaseMockRecorder struct {
	mock *MockAuthUseCase
}

// NewMockAuthUseCase creates a new mock instance.
func NewMockAuthUseCase(ctrl *gomock.Controller) *MockAuthUseCase {
	mock := &MockAuthUseCase{ctrl: ctrl}
	mock.recorder = &MockAuthUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthUseCase) EXPECT() *MockAuthUseCaseMockRecorder {
	return m.recorder
}

// SendOTP mocks base method.
func (m *MockAuthUseCase) SendOTP(phone string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendOTP", phone)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendOTP indicates an expected call of SendOTP.
func (mr *MockAuthUseCaseMockRecorder) SendOTP(phone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendOTP", reflect.TypeOf((*MockAuthUseCase)(nil).SendOTP), phone)
}

// SendVerificationEmail mocks base method.
func (m *MockAuthUseCase) SendVerificationEmail(email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendVerificationEmail", email)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendVerificationEmail indicates an expected call of SendVerificationEmail.
func (mr *MockAuthUseCaseMockRecorder) SendVerificationEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendVerificationEmail", reflect.TypeOf((*MockAuthUseCase)(nil).SendVerificationEmail), email)
}

// UserVerifyAccount mocks base method.
func (m *MockAuthUseCase) UserVerifyAccount(email, code string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserVerifyAccount", email, code)
	ret0, _ := ret[0].(error)
	return ret0
}

// UserVerifyAccount indicates an expected call of UserVerifyAccount.
func (mr *MockAuthUseCaseMockRecorder) UserVerifyAccount(email, code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserVerifyAccount", reflect.TypeOf((*MockAuthUseCase)(nil).UserVerifyAccount), email, code)
}

// VarifyOTP mocks base method.
func (m *MockAuthUseCase) VarifyOTP(phone, otp string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VarifyOTP", phone, otp)
	ret0, _ := ret[0].(error)
	return ret0
}

// VarifyOTP indicates an expected call of VarifyOTP.
func (mr *MockAuthUseCaseMockRecorder) VarifyOTP(phone, otp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VarifyOTP", reflect.TypeOf((*MockAuthUseCase)(nil).VarifyOTP), phone, otp)
}

// VerifyAdmin mocks base method.
func (m *MockAuthUseCase) VerifyAdmin(email, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyAdmin", email, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyAdmin indicates an expected call of VerifyAdmin.
func (mr *MockAuthUseCaseMockRecorder) VerifyAdmin(email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyAdmin", reflect.TypeOf((*MockAuthUseCase)(nil).VerifyAdmin), email, password)
}

// VerifyUser mocks base method.
func (m *MockAuthUseCase) VerifyUser(email, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyUser", email, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyUser indicates an expected call of VerifyUser.
func (mr *MockAuthUseCaseMockRecorder) VerifyUser(email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyUser", reflect.TypeOf((*MockAuthUseCase)(nil).VerifyUser), email, password)
}

// VerifyWorker mocks base method.
func (m *MockAuthUseCase) VerifyWorker(email, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyWorker", email, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyWorker indicates an expected call of VerifyWorker.
func (mr *MockAuthUseCaseMockRecorder) VerifyWorker(email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyWorker", reflect.TypeOf((*MockAuthUseCase)(nil).VerifyWorker), email, password)
}

// WorkerVerifyAccount mocks base method.
func (m *MockAuthUseCase) WorkerVerifyAccount(email string, code int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WorkerVerifyAccount", email, code)
	ret0, _ := ret[0].(error)
	return ret0
}

// WorkerVerifyAccount indicates an expected call of WorkerVerifyAccount.
func (mr *MockAuthUseCaseMockRecorder) WorkerVerifyAccount(email, code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WorkerVerifyAccount", reflect.TypeOf((*MockAuthUseCase)(nil).WorkerVerifyAccount), email, code)
}