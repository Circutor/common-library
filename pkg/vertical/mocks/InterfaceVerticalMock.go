// Code generated by mockery v2.5.1. DO NOT EDIT.

package mocks

import (
	data "github.com/circutor/common-library/pkg/data"
	controller "github.com/circutor/thingsboard-methods/pkg/controller"

	mock "github.com/stretchr/testify/mock"

	request "github.com/circutor/common-library/pkg/request"
)

// InterfaceVerticalMock is an autogenerated mock type for the InterfaceVertical type
type InterfaceVerticalMock struct {
	mock.Mock
}

// GetDeviceAlarms provides a mock function with given fields: deviceID, token, msg, tb, _a4, _a5
func (_m *InterfaceVerticalMock) GetDeviceAlarms(deviceID string, token string, msg string, tb controller.ThingsBoardController, _a4 data.InterfaceData, _a5 request.InterfaceRequest) (int, map[string]interface{}, error) {
	ret := _m.Called(deviceID, token, msg, tb, _a4, _a5)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, string, string, controller.ThingsBoardController, data.InterfaceData, request.InterfaceRequest) int); ok {
		r0 = rf(deviceID, token, msg, tb, _a4, _a5)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 map[string]interface{}
	if rf, ok := ret.Get(1).(func(string, string, string, controller.ThingsBoardController, data.InterfaceData, request.InterfaceRequest) map[string]interface{}); ok {
		r1 = rf(deviceID, token, msg, tb, _a4, _a5)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]interface{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string, string, controller.ThingsBoardController, data.InterfaceData, request.InterfaceRequest) error); ok {
		r2 = rf(deviceID, token, msg, tb, _a4, _a5)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SetAttrClient provides a mock function with given fields: accessToken, msg, tb, _a3
func (_m *InterfaceVerticalMock) SetAttrClient(accessToken string, msg string, tb controller.ThingsBoardController, _a3 data.InterfaceData) (int, map[string]interface{}, error) {
	ret := _m.Called(accessToken, msg, tb, _a3)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, string, controller.ThingsBoardController, data.InterfaceData) int); ok {
		r0 = rf(accessToken, msg, tb, _a3)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 map[string]interface{}
	if rf, ok := ret.Get(1).(func(string, string, controller.ThingsBoardController, data.InterfaceData) map[string]interface{}); ok {
		r1 = rf(accessToken, msg, tb, _a3)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]interface{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string, controller.ThingsBoardController, data.InterfaceData) error); ok {
		r2 = rf(accessToken, msg, tb, _a3)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SetAttrServer provides a mock function with given fields: deviceID, costumerID, secretKey, token, msg, expirationTime, tb, _a7
func (_m *InterfaceVerticalMock) SetAttrServer(deviceID string, costumerID string, secretKey string, token string, msg string, expirationTime int, tb controller.ThingsBoardController, _a7 data.InterfaceData) (int, map[string]interface{}, error) {
	ret := _m.Called(deviceID, costumerID, secretKey, token, msg, expirationTime, tb, _a7)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, string, string, string, string, int, controller.ThingsBoardController, data.InterfaceData) int); ok {
		r0 = rf(deviceID, costumerID, secretKey, token, msg, expirationTime, tb, _a7)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 map[string]interface{}
	if rf, ok := ret.Get(1).(func(string, string, string, string, string, int, controller.ThingsBoardController, data.InterfaceData) map[string]interface{}); ok {
		r1 = rf(deviceID, costumerID, secretKey, token, msg, expirationTime, tb, _a7)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]interface{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string, string, string, string, int, controller.ThingsBoardController, data.InterfaceData) error); ok {
		r2 = rf(deviceID, costumerID, secretKey, token, msg, expirationTime, tb, _a7)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// TelemetryInitialization provides a mock function with given fields: accessToken, deviceType, msg, tb
func (_m *InterfaceVerticalMock) TelemetryInitialization(accessToken string, deviceType string, msg string, tb controller.ThingsBoardController) (int, map[string]interface{}, error) {
	ret := _m.Called(accessToken, deviceType, msg, tb)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, string, string, controller.ThingsBoardController) int); ok {
		r0 = rf(accessToken, deviceType, msg, tb)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 map[string]interface{}
	if rf, ok := ret.Get(1).(func(string, string, string, controller.ThingsBoardController) map[string]interface{}); ok {
		r1 = rf(accessToken, deviceType, msg, tb)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]interface{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string, string, controller.ThingsBoardController) error); ok {
		r2 = rf(accessToken, deviceType, msg, tb)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateTypology provides a mock function with given fields: deviceID, token, msg, tb, _a4
func (_m *InterfaceVerticalMock) UpdateTypology(deviceID string, token string, msg string, tb controller.ThingsBoardController, _a4 data.InterfaceData) (int, map[string]interface{}, error) {
	ret := _m.Called(deviceID, token, msg, tb, _a4)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, string, string, controller.ThingsBoardController, data.InterfaceData) int); ok {
		r0 = rf(deviceID, token, msg, tb, _a4)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 map[string]interface{}
	if rf, ok := ret.Get(1).(func(string, string, string, controller.ThingsBoardController, data.InterfaceData) map[string]interface{}); ok {
		r1 = rf(deviceID, token, msg, tb, _a4)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]interface{})
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, string, string, controller.ThingsBoardController, data.InterfaceData) error); ok {
		r2 = rf(deviceID, token, msg, tb, _a4)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
