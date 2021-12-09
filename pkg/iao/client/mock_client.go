// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/shenghui0779/yiigo (interfaces: HTTPClient,UploadForm)

// Package client is a generated GoMock package.
package client

import (
	context "context"
	multipart "mime/multipart"
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	yiigo "github.com/shenghui0779/yiigo"
)

// MockHTTPClient is a mock of HTTPClient interface.
type MockHTTPClient struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPClientMockRecorder
}

// MockHTTPClientMockRecorder is the mock recorder for MockHTTPClient.
type MockHTTPClientMockRecorder struct {
	mock *MockHTTPClient
}

// NewMockHTTPClient creates a new mock instance.
func NewMockHTTPClient(ctrl *gomock.Controller) *MockHTTPClient {
	mock := &MockHTTPClient{ctrl: ctrl}
	mock.recorder = &MockHTTPClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHTTPClient) EXPECT() *MockHTTPClientMockRecorder {
	return m.recorder
}

// Do mocks base method.
func (m *MockHTTPClient) Do(arg0 context.Context, arg1, arg2 string, arg3 []byte, arg4 ...yiigo.HTTPOption) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Do", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do.
func (mr *MockHTTPClientMockRecorder) Do(arg0, arg1, arg2, arg3 interface{}, arg4 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockHTTPClient)(nil).Do), varargs...)
}

// Upload mocks base method.
func (m *MockHTTPClient) Upload(arg0 context.Context, arg1 string, arg2 yiigo.UploadForm, arg3 ...yiigo.HTTPOption) (*http.Response, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Upload", varargs...)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upload indicates an expected call of Upload.
func (mr *MockHTTPClientMockRecorder) Upload(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockHTTPClient)(nil).Upload), varargs...)
}

// MockUploadForm is a mock of UploadForm interface.
type MockUploadForm struct {
	ctrl     *gomock.Controller
	recorder *MockUploadFormMockRecorder
}

// MockUploadFormMockRecorder is the mock recorder for MockUploadForm.
type MockUploadFormMockRecorder struct {
	mock *MockUploadForm
}

// NewMockUploadForm creates a new mock instance.
func NewMockUploadForm(ctrl *gomock.Controller) *MockUploadForm {
	mock := &MockUploadForm{ctrl: ctrl}
	mock.recorder = &MockUploadFormMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUploadForm) EXPECT() *MockUploadFormMockRecorder {
	return m.recorder
}

// Write mocks base method.
func (m *MockUploadForm) Write(arg0 *multipart.Writer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *MockUploadFormMockRecorder) Write(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockUploadForm)(nil).Write), arg0)
}
