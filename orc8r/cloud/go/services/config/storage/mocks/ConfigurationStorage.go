/*
 * Copyright (c) Facebook, Inc. and its affiliates.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import gostorage "magma/orc8r/cloud/go/storage"
import mock "github.com/stretchr/testify/mock"
import storage "magma/orc8r/cloud/go/services/config/storage"

// ConfigurationStorage is an autogenerated mock type for the ConfigurationStorage type
type ConfigurationStorage struct {
	mock.Mock
}

// CreateConfig provides a mock function with given fields: networkId, configType, key, value
func (_m *ConfigurationStorage) CreateConfig(networkId string, configType string, key string, value []byte) error {
	ret := _m.Called(networkId, configType, key, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, []byte) error); ok {
		r0 = rf(networkId, configType, key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteConfig provides a mock function with given fields: networkId, configType, key
func (_m *ConfigurationStorage) DeleteConfig(networkId string, configType string, key string) error {
	ret := _m.Called(networkId, configType, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(networkId, configType, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteConfigs provides a mock function with given fields: networkId, criteria
func (_m *ConfigurationStorage) DeleteConfigs(networkId string, criteria *storage.FilterCriteria) error {
	ret := _m.Called(networkId, criteria)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *storage.FilterCriteria) error); ok {
		r0 = rf(networkId, criteria)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteConfigsForNetwork provides a mock function with given fields: networkId
func (_m *ConfigurationStorage) DeleteConfigsForNetwork(networkId string) error {
	ret := _m.Called(networkId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(networkId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetConfig provides a mock function with given fields: networkId, configType, key
func (_m *ConfigurationStorage) GetConfig(networkId string, configType string, key string) (*storage.ConfigValue, error) {
	ret := _m.Called(networkId, configType, key)

	var r0 *storage.ConfigValue
	if rf, ok := ret.Get(0).(func(string, string, string) *storage.ConfigValue); ok {
		r0 = rf(networkId, configType, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.ConfigValue)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(networkId, configType, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConfigs provides a mock function with given fields: networkId, criteria
func (_m *ConfigurationStorage) GetConfigs(networkId string, criteria *storage.FilterCriteria) (map[gostorage.TypeAndKey]*storage.ConfigValue, error) {
	ret := _m.Called(networkId, criteria)

	var r0 map[gostorage.TypeAndKey]*storage.ConfigValue
	if rf, ok := ret.Get(0).(func(string, *storage.FilterCriteria) map[gostorage.TypeAndKey]*storage.ConfigValue); ok {
		r0 = rf(networkId, criteria)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[gostorage.TypeAndKey]*storage.ConfigValue)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *storage.FilterCriteria) error); ok {
		r1 = rf(networkId, criteria)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListKeysForType provides a mock function with given fields: networkId, configType
func (_m *ConfigurationStorage) ListKeysForType(networkId string, configType string) ([]string, error) {
	ret := _m.Called(networkId, configType)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string, string) []string); ok {
		r0 = rf(networkId, configType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(networkId, configType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateConfig provides a mock function with given fields: networkId, configType, key, newValue
func (_m *ConfigurationStorage) UpdateConfig(networkId string, configType string, key string, newValue []byte) error {
	ret := _m.Called(networkId, configType, key, newValue)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, []byte) error); ok {
		r0 = rf(networkId, configType, key, newValue)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
