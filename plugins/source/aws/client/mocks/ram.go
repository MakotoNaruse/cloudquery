// Code generated by MockGen. DO NOT EDIT.
// Source: ram.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	ram "github.com/aws/aws-sdk-go-v2/service/ram"
	gomock "github.com/golang/mock/gomock"
)

// MockRamClient is a mock of RamClient interface.
type MockRamClient struct {
	ctrl     *gomock.Controller
	recorder *MockRamClientMockRecorder
}

// MockRamClientMockRecorder is the mock recorder for MockRamClient.
type MockRamClientMockRecorder struct {
	mock *MockRamClient
}

// NewMockRamClient creates a new mock instance.
func NewMockRamClient(ctrl *gomock.Controller) *MockRamClient {
	mock := &MockRamClient{ctrl: ctrl}
	mock.recorder = &MockRamClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRamClient) EXPECT() *MockRamClientMockRecorder {
	return m.recorder
}

// GetPermission mocks base method.
func (m *MockRamClient) GetPermission(arg0 context.Context, arg1 *ram.GetPermissionInput, arg2 ...func(*ram.Options)) (*ram.GetPermissionOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &ram.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetPermission")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPermission", varargs...)
	ret0, _ := ret[0].(*ram.GetPermissionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPermission indicates an expected call of GetPermission.
func (mr *MockRamClientMockRecorder) GetPermission(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPermission", reflect.TypeOf((*MockRamClient)(nil).GetPermission), varargs...)
}

// GetResourcePolicies mocks base method.
func (m *MockRamClient) GetResourcePolicies(arg0 context.Context, arg1 *ram.GetResourcePoliciesInput, arg2 ...func(*ram.Options)) (*ram.GetResourcePoliciesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &ram.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetResourcePolicies")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetResourcePolicies", varargs...)
	ret0, _ := ret[0].(*ram.GetResourcePoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResourcePolicies indicates an expected call of GetResourcePolicies.
func (mr *MockRamClientMockRecorder) GetResourcePolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourcePolicies", reflect.TypeOf((*MockRamClient)(nil).GetResourcePolicies), varargs...)
}

// GetResourceShareAssociations mocks base method.
func (m *MockRamClient) GetResourceShareAssociations(arg0 context.Context, arg1 *ram.GetResourceShareAssociationsInput, arg2 ...func(*ram.Options)) (*ram.GetResourceShareAssociationsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &ram.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetResourceShareAssociations")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetResourceShareAssociations", varargs...)
	ret0, _ := ret[0].(*ram.GetResourceShareAssociationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResourceShareAssociations indicates an expected call of GetResourceShareAssociations.
func (mr *MockRamClientMockRecorder) GetResourceShareAssociations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceShareAssociations", reflect.TypeOf((*MockRamClient)(nil).GetResourceShareAssociations), varargs...)
}

// GetResourceShareInvitations mocks base method.
func (m *MockRamClient) GetResourceShareInvitations(arg0 context.Context, arg1 *ram.GetResourceShareInvitationsInput, arg2 ...func(*ram.Options)) (*ram.GetResourceShareInvitationsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &ram.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetResourceShareInvitations")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetResourceShareInvitations", varargs...)
	ret0, _ := ret[0].(*ram.GetResourceShareInvitationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResourceShareInvitations indicates an expected call of GetResourceShareInvitations.
func (mr *MockRamClientMockRecorder) GetResourceShareInvitations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceShareInvitations", reflect.TypeOf((*MockRamClient)(nil).GetResourceShareInvitations), varargs...)
}

// GetResourceShares mocks base method.
func (m *MockRamClient) GetResourceShares(arg0 context.Context, arg1 *ram.GetResourceSharesInput, arg2 ...func(*ram.Options)) (*ram.GetResourceSharesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &ram.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetResourceShares")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetResourceShares", varargs...)
	ret0, _ := ret[0].(*ram.GetResourceSharesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResourceShares indicates an expected call of GetResourceShares.
func (mr *MockRamClientMockRecorder) GetResourceShares(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceShares", reflect.TypeOf((*MockRamClient)(nil).GetResourceShares), varargs...)
}

// ListPendingInvitationResources mocks base method.
func (m *MockRamClient) ListPendingInvitationResources(arg0 context.Context, arg1 *ram.ListPendingInvitationResourcesInput, arg2 ...func(*ram.Options)) (*ram.ListPendingInvitationResourcesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &ram.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListPendingInvitationResources")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPendingInvitationResources", varargs...)
	ret0, _ := ret[0].(*ram.ListPendingInvitationResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPendingInvitationResources indicates an expected call of ListPendingInvitationResources.
func (mr *MockRamClientMockRecorder) ListPendingInvitationResources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPendingInvitationResources", reflect.TypeOf((*MockRamClient)(nil).ListPendingInvitationResources), varargs...)
}

// ListPermissionVersions mocks base method.
func (m *MockRamClient) ListPermissionVersions(arg0 context.Context, arg1 *ram.ListPermissionVersionsInput, arg2 ...func(*ram.Options)) (*ram.ListPermissionVersionsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &ram.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListPermissionVersions")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPermissionVersions", varargs...)
	ret0, _ := ret[0].(*ram.ListPermissionVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPermissionVersions indicates an expected call of ListPermissionVersions.
func (mr *MockRamClientMockRecorder) ListPermissionVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPermissionVersions", reflect.TypeOf((*MockRamClient)(nil).ListPermissionVersions), varargs...)
}

// ListPermissions mocks base method.
func (m *MockRamClient) ListPermissions(arg0 context.Context, arg1 *ram.ListPermissionsInput, arg2 ...func(*ram.Options)) (*ram.ListPermissionsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &ram.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListPermissions")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPermissions", varargs...)
	ret0, _ := ret[0].(*ram.ListPermissionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPermissions indicates an expected call of ListPermissions.
func (mr *MockRamClientMockRecorder) ListPermissions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPermissions", reflect.TypeOf((*MockRamClient)(nil).ListPermissions), varargs...)
}

// ListPrincipals mocks base method.
func (m *MockRamClient) ListPrincipals(arg0 context.Context, arg1 *ram.ListPrincipalsInput, arg2 ...func(*ram.Options)) (*ram.ListPrincipalsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &ram.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListPrincipals")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPrincipals", varargs...)
	ret0, _ := ret[0].(*ram.ListPrincipalsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPrincipals indicates an expected call of ListPrincipals.
func (mr *MockRamClientMockRecorder) ListPrincipals(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPrincipals", reflect.TypeOf((*MockRamClient)(nil).ListPrincipals), varargs...)
}

// ListResourceSharePermissions mocks base method.
func (m *MockRamClient) ListResourceSharePermissions(arg0 context.Context, arg1 *ram.ListResourceSharePermissionsInput, arg2 ...func(*ram.Options)) (*ram.ListResourceSharePermissionsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &ram.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListResourceSharePermissions")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListResourceSharePermissions", varargs...)
	ret0, _ := ret[0].(*ram.ListResourceSharePermissionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListResourceSharePermissions indicates an expected call of ListResourceSharePermissions.
func (mr *MockRamClientMockRecorder) ListResourceSharePermissions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResourceSharePermissions", reflect.TypeOf((*MockRamClient)(nil).ListResourceSharePermissions), varargs...)
}

// ListResourceTypes mocks base method.
func (m *MockRamClient) ListResourceTypes(arg0 context.Context, arg1 *ram.ListResourceTypesInput, arg2 ...func(*ram.Options)) (*ram.ListResourceTypesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &ram.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListResourceTypes")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListResourceTypes", varargs...)
	ret0, _ := ret[0].(*ram.ListResourceTypesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListResourceTypes indicates an expected call of ListResourceTypes.
func (mr *MockRamClientMockRecorder) ListResourceTypes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResourceTypes", reflect.TypeOf((*MockRamClient)(nil).ListResourceTypes), varargs...)
}

// ListResources mocks base method.
func (m *MockRamClient) ListResources(arg0 context.Context, arg1 *ram.ListResourcesInput, arg2 ...func(*ram.Options)) (*ram.ListResourcesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &ram.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListResources")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListResources", varargs...)
	ret0, _ := ret[0].(*ram.ListResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListResources indicates an expected call of ListResources.
func (mr *MockRamClientMockRecorder) ListResources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResources", reflect.TypeOf((*MockRamClient)(nil).ListResources), varargs...)
}
