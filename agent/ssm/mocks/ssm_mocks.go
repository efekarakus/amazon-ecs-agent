// Copyright 2015-2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-ecs-agent/agent/ssm (interfaces: SSMClient)

// Package mock_ssm is a generated GoMock package.
package mock_ssm

import (
	reflect "reflect"

	ssm "github.com/aws/aws-sdk-go/service/ssm"
	gomock "github.com/golang/mock/gomock"
)

// MockSSMClient is a mock of SSMClient interface
type MockSSMClient struct {
	ctrl     *gomock.Controller
	recorder *MockSSMClientMockRecorder
}

// MockSSMClientMockRecorder is the mock recorder for MockSSMClient
type MockSSMClientMockRecorder struct {
	mock *MockSSMClient
}

// NewMockSSMClient creates a new mock instance
func NewMockSSMClient(ctrl *gomock.Controller) *MockSSMClient {
	mock := &MockSSMClient{ctrl: ctrl}
	mock.recorder = &MockSSMClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSSMClient) EXPECT() *MockSSMClientMockRecorder {
	return m.recorder
}

// GetParameters mocks base method
func (m *MockSSMClient) GetParameters(arg0 *ssm.GetParametersInput) (*ssm.GetParametersOutput, error) {
	ret := m.ctrl.Call(m, "GetParameters", arg0)
	ret0, _ := ret[0].(*ssm.GetParametersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetParameters indicates an expected call of GetParameters
func (mr *MockSSMClientMockRecorder) GetParameters(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParameters", reflect.TypeOf((*MockSSMClient)(nil).GetParameters), arg0)
}
