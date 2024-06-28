// Code generated by MockGen. DO NOT EDIT.
// Source: types.go

// Package channel is a generated GoMock package.
package channel

import (
	context "context"
	reflect "reflect"

	pgx "github.com/assignment-amori/pkg/sql/pgx"
	v5 "github.com/jackc/pgx/v5"
	pgconn "github.com/jackc/pgx/v5/pgconn"
	gomock "go.uber.org/mock/gomock"
)

// MockdatabaseResource is a mock of databaseResource interface.
type MockdatabaseResource struct {
	ctrl     *gomock.Controller
	recorder *MockdatabaseResourceMockRecorder
}

// MockdatabaseResourceMockRecorder is the mock recorder for MockdatabaseResource.
type MockdatabaseResourceMockRecorder struct {
	mock *MockdatabaseResource
}

// NewMockdatabaseResource creates a new mock instance.
func NewMockdatabaseResource(ctrl *gomock.Controller) *MockdatabaseResource {
	mock := &MockdatabaseResource{ctrl: ctrl}
	mock.recorder = &MockdatabaseResourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockdatabaseResource) EXPECT() *MockdatabaseResourceMockRecorder {
	return m.recorder
}

// ExecuteInTx mocks base method.
func (m *MockdatabaseResource) ExecuteInTx(ctx context.Context, txConsistency *pgx.Tx, fn func(*pgx.Tx) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteInTx", ctx, txConsistency, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExecuteInTx indicates an expected call of ExecuteInTx.
func (mr *MockdatabaseResourceMockRecorder) ExecuteInTx(ctx, txConsistency, fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteInTx", reflect.TypeOf((*MockdatabaseResource)(nil).ExecuteInTx), ctx, txConsistency, fn)
}

// Insert mocks base method.
func (m *MockdatabaseResource) Insert(ctx context.Context, entity interface{}) (pgconn.CommandTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, entity)
	ret0, _ := ret[0].(pgconn.CommandTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockdatabaseResourceMockRecorder) Insert(ctx, entity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockdatabaseResource)(nil).Insert), ctx, entity)
}

// Query mocks base method.
func (m *MockdatabaseResource) Query(ctx context.Context, sql string, args ...any) (v5.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, sql}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(v5.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockdatabaseResourceMockRecorder) Query(ctx, sql interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, sql}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockdatabaseResource)(nil).Query), varargs...)
}

// QueryRow mocks base method.
func (m *MockdatabaseResource) QueryRow(ctx context.Context, sql string, args ...any) v5.Row {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, sql}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryRow", varargs...)
	ret0, _ := ret[0].(v5.Row)
	return ret0
}

// QueryRow indicates an expected call of QueryRow.
func (mr *MockdatabaseResourceMockRecorder) QueryRow(ctx, sql interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, sql}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRow", reflect.TypeOf((*MockdatabaseResource)(nil).QueryRow), varargs...)
}

// Select mocks base method.
func (m *MockdatabaseResource) Select(ctx context.Context, dest interface{}, sql string, args ...interface{}) (bool, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, dest, sql}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Select", varargs...)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Select indicates an expected call of Select.
func (mr *MockdatabaseResourceMockRecorder) Select(ctx, dest, sql interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, dest, sql}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Select", reflect.TypeOf((*MockdatabaseResource)(nil).Select), varargs...)
}

// Update mocks base method.
func (m *MockdatabaseResource) Update(ctx context.Context, entity interface{}, options *pgx.UpdateOptions) (pgconn.CommandTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, entity, options)
	ret0, _ := ret[0].(pgconn.CommandTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockdatabaseResourceMockRecorder) Update(ctx, entity, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockdatabaseResource)(nil).Update), ctx, entity, options)
}

// MocktxResource is a mock of txResource interface.
type MocktxResource struct {
	ctrl     *gomock.Controller
	recorder *MocktxResourceMockRecorder
}

// MocktxResourceMockRecorder is the mock recorder for MocktxResource.
type MocktxResourceMockRecorder struct {
	mock *MocktxResource
}

// NewMocktxResource creates a new mock instance.
func NewMocktxResource(ctrl *gomock.Controller) *MocktxResource {
	mock := &MocktxResource{ctrl: ctrl}
	mock.recorder = &MocktxResourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocktxResource) EXPECT() *MocktxResourceMockRecorder {
	return m.recorder
}

// Beginx mocks base method.
func (m *MocktxResource) Beginx(ctx context.Context) (*pgx.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Beginx", ctx)
	ret0, _ := ret[0].(*pgx.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Beginx indicates an expected call of Beginx.
func (mr *MocktxResourceMockRecorder) Beginx(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Beginx", reflect.TypeOf((*MocktxResource)(nil).Beginx), ctx)
}

// Insert mocks base method.
func (m *MocktxResource) Insert(ctx context.Context, entity interface{}) (pgconn.CommandTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, entity)
	ret0, _ := ret[0].(pgconn.CommandTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MocktxResourceMockRecorder) Insert(ctx, entity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MocktxResource)(nil).Insert), ctx, entity)
}

// Query mocks base method.
func (m *MocktxResource) Query(ctx context.Context, sql string, args ...any) (v5.Rows, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, sql}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(v5.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MocktxResourceMockRecorder) Query(ctx, sql interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, sql}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MocktxResource)(nil).Query), varargs...)
}

// QueryRow mocks base method.
func (m *MocktxResource) QueryRow(ctx context.Context, sql string, args ...any) v5.Row {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, sql}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryRow", varargs...)
	ret0, _ := ret[0].(v5.Row)
	return ret0
}

// QueryRow indicates an expected call of QueryRow.
func (mr *MocktxResourceMockRecorder) QueryRow(ctx, sql interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, sql}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRow", reflect.TypeOf((*MocktxResource)(nil).QueryRow), varargs...)
}

// Update mocks base method.
func (m *MocktxResource) Update(ctx context.Context, entity interface{}, options *pgx.UpdateOptions) (pgconn.CommandTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, entity, options)
	ret0, _ := ret[0].(pgconn.CommandTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MocktxResourceMockRecorder) Update(ctx, entity, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MocktxResource)(nil).Update), ctx, entity, options)
}

// MocksonyFlakeResource is a mock of sonyFlakeResource interface.
type MocksonyFlakeResource struct {
	ctrl     *gomock.Controller
	recorder *MocksonyFlakeResourceMockRecorder
}

// MocksonyFlakeResourceMockRecorder is the mock recorder for MocksonyFlakeResource.
type MocksonyFlakeResourceMockRecorder struct {
	mock *MocksonyFlakeResource
}

// NewMocksonyFlakeResource creates a new mock instance.
func NewMocksonyFlakeResource(ctrl *gomock.Controller) *MocksonyFlakeResource {
	mock := &MocksonyFlakeResource{ctrl: ctrl}
	mock.recorder = &MocksonyFlakeResourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocksonyFlakeResource) EXPECT() *MocksonyFlakeResourceMockRecorder {
	return m.recorder
}

// NextID mocks base method.
func (m *MocksonyFlakeResource) NextID() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextID")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NextID indicates an expected call of NextID.
func (mr *MocksonyFlakeResourceMockRecorder) NextID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextID", reflect.TypeOf((*MocksonyFlakeResource)(nil).NextID))
}
