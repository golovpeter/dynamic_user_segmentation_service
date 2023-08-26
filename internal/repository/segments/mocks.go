// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package segments is a generated GoMock package.
package segments

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreateSegment mocks base method.
func (m *MockRepository) CreateSegment(slug string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSegment", slug)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSegment indicates an expected call of CreateSegment.
func (mr *MockRepositoryMockRecorder) CreateSegment(slug interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSegment", reflect.TypeOf((*MockRepository)(nil).CreateSegment), slug)
}

// DeleteSegment mocks base method.
func (m *MockRepository) DeleteSegment(slug string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSegment", slug)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteSegment indicates an expected call of DeleteSegment.
func (mr *MockRepositoryMockRecorder) DeleteSegment(slug interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSegment", reflect.TypeOf((*MockRepository)(nil).DeleteSegment), slug)
}

// GetActiveSegmentsIdsBySlugs mocks base method.
func (m *MockRepository) GetActiveSegmentsIdsBySlugs(slugs []string) (map[string]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActiveSegmentsIdsBySlugs", slugs)
	ret0, _ := ret[0].(map[string]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActiveSegmentsIdsBySlugs indicates an expected call of GetActiveSegmentsIdsBySlugs.
func (mr *MockRepositoryMockRecorder) GetActiveSegmentsIdsBySlugs(slugs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActiveSegmentsIdsBySlugs", reflect.TypeOf((*MockRepository)(nil).GetActiveSegmentsIdsBySlugs), slugs)
}

// GetUserSegments mocks base method.
func (m *MockRepository) GetUserSegments(id int64) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserSegments", id)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserSegments indicates an expected call of GetUserSegments.
func (mr *MockRepositoryMockRecorder) GetUserSegments(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserSegments", reflect.TypeOf((*MockRepository)(nil).GetUserSegments), id)
}
