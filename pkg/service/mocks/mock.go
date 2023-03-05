// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	entity "github.com/Phaseant/MusicAPI/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockAutorization is a mock of Autorization interface.
type MockAutorization struct {
	ctrl     *gomock.Controller
	recorder *MockAutorizationMockRecorder
}

// MockAutorizationMockRecorder is the mock recorder for MockAutorization.
type MockAutorizationMockRecorder struct {
	mock *MockAutorization
}

// NewMockAutorization creates a new mock instance.
func NewMockAutorization(ctrl *gomock.Controller) *MockAutorization {
	mock := &MockAutorization{ctrl: ctrl}
	mock.recorder = &MockAutorizationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAutorization) EXPECT() *MockAutorizationMockRecorder {
	return m.recorder
}

// GenerateToken mocks base method.
func (m *MockAutorization) GenerateToken(username, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", username, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockAutorizationMockRecorder) GenerateToken(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockAutorization)(nil).GenerateToken), username, password)
}

// NewUser mocks base method.
func (m *MockAutorization) NewUser(user entity.User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUser", user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewUser indicates an expected call of NewUser.
func (mr *MockAutorizationMockRecorder) NewUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUser", reflect.TypeOf((*MockAutorization)(nil).NewUser), user)
}

// ParseToken mocks base method.
func (m *MockAutorization) ParseToken(accessToken string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", accessToken)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockAutorizationMockRecorder) ParseToken(accessToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockAutorization)(nil).ParseToken), accessToken)
}

// MockAlbum is a mock of Album interface.
type MockAlbum struct {
	ctrl     *gomock.Controller
	recorder *MockAlbumMockRecorder
}

// MockAlbumMockRecorder is the mock recorder for MockAlbum.
type MockAlbumMockRecorder struct {
	mock *MockAlbum
}

// NewMockAlbum creates a new mock instance.
func NewMockAlbum(ctrl *gomock.Controller) *MockAlbum {
	mock := &MockAlbum{ctrl: ctrl}
	mock.recorder = &MockAlbumMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAlbum) EXPECT() *MockAlbumMockRecorder {
	return m.recorder
}

// DeleteAlbum mocks base method.
func (m *MockAlbum) DeleteAlbum(id string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAlbum", id)
	ret0, _ := ret[0].(bool)
	return ret0
}

// DeleteAlbum indicates an expected call of DeleteAlbum.
func (mr *MockAlbumMockRecorder) DeleteAlbum(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAlbum", reflect.TypeOf((*MockAlbum)(nil).DeleteAlbum), id)
}

// GetAlbum mocks base method.
func (m *MockAlbum) GetAlbum(id string) (entity.Album, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAlbum", id)
	ret0, _ := ret[0].(entity.Album)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAlbum indicates an expected call of GetAlbum.
func (mr *MockAlbumMockRecorder) GetAlbum(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAlbum", reflect.TypeOf((*MockAlbum)(nil).GetAlbum), id)
}

// GetAllAlbums mocks base method.
func (m *MockAlbum) GetAllAlbums() ([]entity.Album, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllAlbums")
	ret0, _ := ret[0].([]entity.Album)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllAlbums indicates an expected call of GetAllAlbums.
func (mr *MockAlbumMockRecorder) GetAllAlbums() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllAlbums", reflect.TypeOf((*MockAlbum)(nil).GetAllAlbums))
}

// NewAlbum mocks base method.
func (m *MockAlbum) NewAlbum(album entity.Album) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAlbum", album)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewAlbum indicates an expected call of NewAlbum.
func (mr *MockAlbumMockRecorder) NewAlbum(album interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAlbum", reflect.TypeOf((*MockAlbum)(nil).NewAlbum), album)
}

// MockAdmin is a mock of Admin interface.
type MockAdmin struct {
	ctrl     *gomock.Controller
	recorder *MockAdminMockRecorder
}

// MockAdminMockRecorder is the mock recorder for MockAdmin.
type MockAdminMockRecorder struct {
	mock *MockAdmin
}

// NewMockAdmin creates a new mock instance.
func NewMockAdmin(ctrl *gomock.Controller) *MockAdmin {
	mock := &MockAdmin{ctrl: ctrl}
	mock.recorder = &MockAdminMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdmin) EXPECT() *MockAdminMockRecorder {
	return m.recorder
}

// AddAdmin mocks base method.
func (m *MockAdmin) AddAdmin(admin entity.Admin) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAdmin", admin)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAdmin indicates an expected call of AddAdmin.
func (mr *MockAdminMockRecorder) AddAdmin(admin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAdmin", reflect.TypeOf((*MockAdmin)(nil).AddAdmin), admin)
}

// IsAdmin mocks base method.
func (m *MockAdmin) IsAdmin(id string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAdmin", id)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsAdmin indicates an expected call of IsAdmin.
func (mr *MockAdminMockRecorder) IsAdmin(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAdmin", reflect.TypeOf((*MockAdmin)(nil).IsAdmin), id)
}