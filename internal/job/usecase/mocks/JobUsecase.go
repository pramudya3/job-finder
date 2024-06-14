// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"
	model "job-finder/internal/job/model"

	mock "github.com/stretchr/testify/mock"
)

// JobUsecase is an autogenerated mock type for the JobUsecase type
type JobUsecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, job
func (_m *JobUsecase) Create(ctx context.Context, job *model.CreateJobReq) error {
	ret := _m.Called(ctx, job)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.CreateJobReq) error); ok {
		r0 = rf(ctx, job)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindJob provides a mock function with given fields: ctx, query
func (_m *JobUsecase) FindJob(ctx context.Context, query map[string]interface{}) ([]*model.Job, interface{}, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for FindJob")
	}

	var r0 []*model.Job
	var r1 interface{}
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, map[string]interface{}) ([]*model.Job, interface{}, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, map[string]interface{}) []*model.Job); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Job)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, map[string]interface{}) interface{}); ok {
		r1 = rf(ctx, query)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(interface{})
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, map[string]interface{}) error); ok {
		r2 = rf(ctx, query)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewJobUsecase creates a new instance of JobUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewJobUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *JobUsecase {
	mock := &JobUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
