// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	state "github.com/tendermint/tendermint/state"

	types "github.com/tendermint/tendermint/types"
)

// EvidencePool is an autogenerated mock type for the EvidencePool type
type EvidencePool struct {
	mock.Mock
}

// CheckEvidence provides a mock function with given fields: _a0
func (_m *EvidencePool) CheckEvidence(_a0 types.EvidenceList) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.EvidenceList) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PendingEvidence provides a mock function with given fields: _a0
func (_m *EvidencePool) PendingEvidence(_a0 uint32) []types.Evidence {
	ret := _m.Called(_a0)

	var r0 []types.Evidence
	if rf, ok := ret.Get(0).(func(uint32) []types.Evidence); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Evidence)
		}
	}

	return r0
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *EvidencePool) Update(_a0 *types.Block, _a1 state.State) {
	_m.Called(_a0, _a1)
}
