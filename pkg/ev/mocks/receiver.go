// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	ev "github.com/facebookincubator/symphony/pkg/ev"
	mock "github.com/stretchr/testify/mock"
)

// Receiver is an autogenerated mock type for the Receiver type
type Receiver struct {
	mock.Mock
}

// Receive provides a mock function with given fields: _a0
func (_m *Receiver) Receive(_a0 context.Context) (*ev.Event, error) {
	ret := _m.Called(_a0)

	var r0 *ev.Event
	if rf, ok := ret.Get(0).(func(context.Context) *ev.Event); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ev.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Shutdown provides a mock function with given fields: _a0
func (_m *Receiver) Shutdown(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}