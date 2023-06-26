// Code generated by MockGen. DO NOT EDIT.
// Source: ./user_room.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/kod-source/docker-goa-next/app/model"
)

// MockUserRoomRepository is a mock of UserRoomRepository interface.
type MockUserRoomRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRoomRepositoryMockRecorder
}

// MockUserRoomRepositoryMockRecorder is the mock recorder for MockUserRoomRepository.
type MockUserRoomRepositoryMockRecorder struct {
	mock *MockUserRoomRepository
}

// NewMockUserRoomRepository creates a new mock instance.
func NewMockUserRoomRepository(ctrl *gomock.Controller) *MockUserRoomRepository {
	mock := &MockUserRoomRepository{ctrl: ctrl}
	mock.recorder = &MockUserRoomRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRoomRepository) EXPECT() *MockUserRoomRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserRoomRepository) Create(ctx context.Context, roomID model.RoomID, userID model.UserID) (*model.UserRoom, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, roomID, userID)
	ret0, _ := ret[0].(*model.UserRoom)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserRoomRepositoryMockRecorder) Create(ctx, roomID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserRoomRepository)(nil).Create), ctx, roomID, userID)
}

// Delete mocks base method.
func (m *MockUserRoomRepository) Delete(ctx context.Context, id model.UserRoomID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUserRoomRepositoryMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserRoomRepository)(nil).Delete), ctx, id)
}
