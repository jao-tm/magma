// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import reindex "magma/orc8r/cloud/go/services/state/indexer/reindex"

// Reindexer is an autogenerated mock type for the Reindexer type
type Reindexer struct {
	mock.Mock
}

// GetIndexerVersions provides a mock function with given fields:
func (_m *Reindexer) GetIndexerVersions() ([]*reindex.Version, error) {
	ret := _m.Called()

	var r0 []*reindex.Version
	if rf, ok := ret.Get(0).(func() []*reindex.Version); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*reindex.Version)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Run provides a mock function with given fields: ctx
func (_m *Reindexer) Run(ctx context.Context) {
	_m.Called(ctx)
}

// RunUnsafe provides a mock function with given fields: ctx, indexerID, sendUpdate
func (_m *Reindexer) RunUnsafe(ctx context.Context, indexerID string, sendUpdate func(string)) error {
	ret := _m.Called(ctx, indexerID, sendUpdate)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, func(string)) error); ok {
		r0 = rf(ctx, indexerID, sendUpdate)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
