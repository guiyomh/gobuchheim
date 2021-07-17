// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	buchheim "github.com/guiyomh/gobuchheim/pkg/buchheim"
	mock "github.com/stretchr/testify/mock"
)

// Node is an autogenerated mock type for the Node type
type Node struct {
	mock.Mock
}

// AddIncomingLink provides a mock function with given fields: _a0
func (_m *Node) AddIncomingLink(_a0 buchheim.Link) {
	_m.Called(_a0)
}

// AddOutgoingLink provides a mock function with given fields: _a0
func (_m *Node) AddOutgoingLink(_a0 buchheim.Link) {
	_m.Called(_a0)
}

// ID provides a mock function with given fields:
func (_m *Node) ID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// IDX provides a mock function with given fields:
func (_m *Node) IDX() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// IncomingLink provides a mock function with given fields:
func (_m *Node) IncomingLink() buchheim.LinkList {
	ret := _m.Called()

	var r0 buchheim.LinkList
	if rf, ok := ret.Get(0).(func() buchheim.LinkList); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(buchheim.LinkList)
		}
	}

	return r0
}

// Label provides a mock function with given fields:
func (_m *Node) Label() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// OutgoingLink provides a mock function with given fields:
func (_m *Node) OutgoingLink() buchheim.LinkList {
	ret := _m.Called()

	var r0 buchheim.LinkList
	if rf, ok := ret.Get(0).(func() buchheim.LinkList); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(buchheim.LinkList)
		}
	}

	return r0
}

// SetIDX provides a mock function with given fields: _a0
func (_m *Node) SetIDX(_a0 int) {
	_m.Called(_a0)
}

// SetX provides a mock function with given fields: _a0
func (_m *Node) SetX(_a0 float64) {
	_m.Called(_a0)
}

// SetY provides a mock function with given fields: _a0
func (_m *Node) SetY(_a0 float64) {
	_m.Called(_a0)
}

// String provides a mock function with given fields:
func (_m *Node) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// X provides a mock function with given fields:
func (_m *Node) X() float64 {
	ret := _m.Called()

	var r0 float64
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// Y provides a mock function with given fields:
func (_m *Node) Y() float64 {
	ret := _m.Called()

	var r0 float64
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}
