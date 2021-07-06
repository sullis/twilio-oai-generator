// Code generated by MockGen. DO NOT EDIT.
// Source: client/provider.go

// Package unit is a generated GoMock package.
package unit

import (
	reflect "reflect"
	openapi "twilio-oai-generator/go/rest/api/v2010"

	gomock "github.com/golang/mock/gomock"
)

// MockApiV2010 is a mock of ApiV2010 interface.
type MockApiV2010 struct {
	ctrl     *gomock.Controller
	recorder *MockApiV2010MockRecorder
}

// MockApiV2010MockRecorder is the mock recorder for MockApiV2010.
type MockApiV2010MockRecorder struct {
	mock *MockApiV2010
}

// NewMockApiV2010 creates a new mock instance.
func NewMockApiV2010(ctrl *gomock.Controller) *MockApiV2010 {
	mock := &MockApiV2010{ctrl: ctrl}
	mock.recorder = &MockApiV2010MockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApiV2010) EXPECT() *MockApiV2010MockRecorder {
	return m.recorder
}

// CreateCallRecording mocks base method.
func (m *MockApiV2010) CreateCallRecording(CallSid string, params *openapi.CreateCallRecordingParams) (*openapi.ApiV2010AccountCallCallRecording, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCallRecording", CallSid, params)
	ret0, _ := ret[0].(*openapi.ApiV2010AccountCallCallRecording)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCallRecording indicates an expected call of CreateCallRecording.
func (mr *MockApiV2010MockRecorder) CreateCallRecording(CallSid, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCallRecording", reflect.TypeOf((*MockApiV2010)(nil).CreateCallRecording), CallSid, params)
}

// CreateCredentialAws mocks base method.
func (m *MockApiV2010) CreateCredentialAws(params *openapi.CreateCredentialAwsParams) (*openapi.AccountsV1CredentialCredentialAws, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCredentialAws", params)
	ret0, _ := ret[0].(*openapi.AccountsV1CredentialCredentialAws)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCredentialAws indicates an expected call of CreateCredentialAws.
func (mr *MockApiV2010MockRecorder) CreateCredentialAws(params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCredentialAws", reflect.TypeOf((*MockApiV2010)(nil).CreateCredentialAws), params)
}

// DeleteCallRecording mocks base method.
func (m *MockApiV2010) DeleteCallRecording(CallSid string, Sid int, params *openapi.DeleteCallRecordingParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCallRecording", CallSid, Sid, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCallRecording indicates an expected call of DeleteCallRecording.
func (mr *MockApiV2010MockRecorder) DeleteCallRecording(CallSid, Sid, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCallRecording", reflect.TypeOf((*MockApiV2010)(nil).DeleteCallRecording), CallSid, Sid, params)
}

// DeleteCredentialAws mocks base method.
func (m *MockApiV2010) DeleteCredentialAws(Sid string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCredentialAws", Sid)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCredentialAws indicates an expected call of DeleteCredentialAws.
func (mr *MockApiV2010MockRecorder) DeleteCredentialAws(Sid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCredentialAws", reflect.TypeOf((*MockApiV2010)(nil).DeleteCredentialAws), Sid)
}

// FetchCallRecording mocks base method.
func (m *MockApiV2010) FetchCallRecording(CallSid string, Sid int, params *openapi.FetchCallRecordingParams) (*openapi.ApiV2010AccountCallCallRecording, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchCallRecording", CallSid, Sid, params)
	ret0, _ := ret[0].(*openapi.ApiV2010AccountCallCallRecording)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchCallRecording indicates an expected call of FetchCallRecording.
func (mr *MockApiV2010MockRecorder) FetchCallRecording(CallSid, Sid, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchCallRecording", reflect.TypeOf((*MockApiV2010)(nil).FetchCallRecording), CallSid, Sid, params)
}

// FetchCredentialAws mocks base method.
func (m *MockApiV2010) FetchCredentialAws(Sid string) (*openapi.AccountsV1CredentialCredentialAws, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchCredentialAws", Sid)
	ret0, _ := ret[0].(*openapi.AccountsV1CredentialCredentialAws)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchCredentialAws indicates an expected call of FetchCredentialAws.
func (mr *MockApiV2010MockRecorder) FetchCredentialAws(Sid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchCredentialAws", reflect.TypeOf((*MockApiV2010)(nil).FetchCredentialAws), Sid)
}

// UpdateCallRecording mocks base method.
func (m *MockApiV2010) UpdateCallRecording(CallSid string, Sid int, params *openapi.UpdateCallRecordingParams) (*openapi.ApiV2010AccountCallCallRecording, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCallRecording", CallSid, Sid, params)
	ret0, _ := ret[0].(*openapi.ApiV2010AccountCallCallRecording)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCallRecording indicates an expected call of UpdateCallRecording.
func (mr *MockApiV2010MockRecorder) UpdateCallRecording(CallSid, Sid, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCallRecording", reflect.TypeOf((*MockApiV2010)(nil).UpdateCallRecording), CallSid, Sid, params)
}

// UpdateCredentialAws mocks base method.
func (m *MockApiV2010) UpdateCredentialAws(Sid string, params *openapi.UpdateCredentialAwsParams) (*openapi.AccountsV1CredentialCredentialAws, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCredentialAws", Sid, params)
	ret0, _ := ret[0].(*openapi.AccountsV1CredentialCredentialAws)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCredentialAws indicates an expected call of UpdateCredentialAws.
func (mr *MockApiV2010MockRecorder) UpdateCredentialAws(Sid, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCredentialAws", reflect.TypeOf((*MockApiV2010)(nil).UpdateCredentialAws), Sid, params)
}