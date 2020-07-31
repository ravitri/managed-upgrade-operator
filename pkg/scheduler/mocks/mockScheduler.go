// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/managed-upgrade-operator/pkg/scheduler (interfaces: Scheduler)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/openshift/managed-upgrade-operator/pkg/apis/upgrade/v1alpha1"
	metrics "github.com/openshift/managed-upgrade-operator/pkg/metrics"
	reflect "reflect"
	time "time"
)

// MockScheduler is a mock of Scheduler interface
type MockScheduler struct {
	ctrl     *gomock.Controller
	recorder *MockSchedulerMockRecorder
}

// MockSchedulerMockRecorder is the mock recorder for MockScheduler
type MockSchedulerMockRecorder struct {
	mock *MockScheduler
}

// NewMockScheduler creates a new mock instance
func NewMockScheduler(ctrl *gomock.Controller) *MockScheduler {
	mock := &MockScheduler{ctrl: ctrl}
	mock.recorder = &MockSchedulerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockScheduler) EXPECT() *MockSchedulerMockRecorder {
	return m.recorder
}

// IsReadyToUpgrade mocks base method
func (m *MockScheduler) IsReadyToUpgrade(arg0 *v1alpha1.UpgradeConfig, arg1 metrics.Metrics, arg2 time.Duration) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsReadyToUpgrade", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsReadyToUpgrade indicates an expected call of IsReadyToUpgrade
func (mr *MockSchedulerMockRecorder) IsReadyToUpgrade(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsReadyToUpgrade", reflect.TypeOf((*MockScheduler)(nil).IsReadyToUpgrade), arg0, arg1, arg2)
}