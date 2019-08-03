// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardenctl/pkg/cmd (interfaces: ConfigReader)

// Package cmd is a generated GoMock package.
package cmd

import (
	cmd "github.com/gardener/gardenctl/pkg/cmd"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockConfigReader is a mock of ConfigReader interface
type MockConfigReader struct {
	ctrl     *gomock.Controller
	recorder *MockConfigReaderMockRecorder
}

// MockConfigReaderMockRecorder is the mock recorder for MockConfigReader
type MockConfigReaderMockRecorder struct {
	mock *MockConfigReader
}

// NewMockConfigReader creates a new mock instance
func NewMockConfigReader(ctrl *gomock.Controller) *MockConfigReader {
	mock := &MockConfigReader{ctrl: ctrl}
	mock.recorder = &MockConfigReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConfigReader) EXPECT() *MockConfigReaderMockRecorder {
	return m.recorder
}

// ReadConfig mocks base method
func (m *MockConfigReader) ReadConfig(arg0 string) *cmd.GardenConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadConfig", arg0)
	ret0, _ := ret[0].(*cmd.GardenConfig)
	return ret0
}

// ReadConfig indicates an expected call of ReadConfig
func (mr *MockConfigReaderMockRecorder) ReadConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadConfig", reflect.TypeOf((*MockConfigReader)(nil).ReadConfig), arg0)
}
