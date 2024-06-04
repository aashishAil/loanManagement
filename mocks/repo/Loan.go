// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"
	databasemodel "loanManagement/database/model"

	mock "github.com/stretchr/testify/mock"

	model "loanManagement/repo/model"
)

// Loan is an autogenerated mock type for the Loan type
type Loan struct {
	mock.Mock
}

type Loan_Expecter struct {
	mock *mock.Mock
}

func (_m *Loan) EXPECT() *Loan_Expecter {
	return &Loan_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, data
func (_m *Loan) Create(ctx context.Context, data model.CreateLoanInput) (*databasemodel.Loan, error) {
	ret := _m.Called(ctx, data)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *databasemodel.Loan
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.CreateLoanInput) (*databasemodel.Loan, error)); ok {
		return rf(ctx, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.CreateLoanInput) *databasemodel.Loan); ok {
		r0 = rf(ctx, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*databasemodel.Loan)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.CreateLoanInput) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Loan_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type Loan_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - data model.CreateLoanInput
func (_e *Loan_Expecter) Create(ctx interface{}, data interface{}) *Loan_Create_Call {
	return &Loan_Create_Call{Call: _e.mock.On("Create", ctx, data)}
}

func (_c *Loan_Create_Call) Run(run func(ctx context.Context, data model.CreateLoanInput)) *Loan_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.CreateLoanInput))
	})
	return _c
}

func (_c *Loan_Create_Call) Return(_a0 *databasemodel.Loan, _a1 error) *Loan_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Loan_Create_Call) RunAndReturn(run func(context.Context, model.CreateLoanInput) (*databasemodel.Loan, error)) *Loan_Create_Call {
	_c.Call.Return(run)
	return _c
}

// FindAll provides a mock function with given fields: ctx, data
func (_m *Loan) FindAll(ctx context.Context, data model.FindAllLoanInput) ([]*databasemodel.Loan, error) {
	ret := _m.Called(ctx, data)

	if len(ret) == 0 {
		panic("no return value specified for FindAll")
	}

	var r0 []*databasemodel.Loan
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.FindAllLoanInput) ([]*databasemodel.Loan, error)); ok {
		return rf(ctx, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.FindAllLoanInput) []*databasemodel.Loan); ok {
		r0 = rf(ctx, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*databasemodel.Loan)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.FindAllLoanInput) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Loan_FindAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAll'
type Loan_FindAll_Call struct {
	*mock.Call
}

// FindAll is a helper method to define mock.On call
//   - ctx context.Context
//   - data model.FindAllLoanInput
func (_e *Loan_Expecter) FindAll(ctx interface{}, data interface{}) *Loan_FindAll_Call {
	return &Loan_FindAll_Call{Call: _e.mock.On("FindAll", ctx, data)}
}

func (_c *Loan_FindAll_Call) Run(run func(ctx context.Context, data model.FindAllLoanInput)) *Loan_FindAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.FindAllLoanInput))
	})
	return _c
}

func (_c *Loan_FindAll_Call) Return(_a0 []*databasemodel.Loan, _a1 error) *Loan_FindAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Loan_FindAll_Call) RunAndReturn(run func(context.Context, model.FindAllLoanInput) ([]*databasemodel.Loan, error)) *Loan_FindAll_Call {
	_c.Call.Return(run)
	return _c
}

// FindOne provides a mock function with given fields: ctx, data
func (_m *Loan) FindOne(ctx context.Context, data model.FindOneLoanInput) (*databasemodel.Loan, error) {
	ret := _m.Called(ctx, data)

	if len(ret) == 0 {
		panic("no return value specified for FindOne")
	}

	var r0 *databasemodel.Loan
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.FindOneLoanInput) (*databasemodel.Loan, error)); ok {
		return rf(ctx, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.FindOneLoanInput) *databasemodel.Loan); ok {
		r0 = rf(ctx, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*databasemodel.Loan)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.FindOneLoanInput) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Loan_FindOne_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindOne'
type Loan_FindOne_Call struct {
	*mock.Call
}

// FindOne is a helper method to define mock.On call
//   - ctx context.Context
//   - data model.FindOneLoanInput
func (_e *Loan_Expecter) FindOne(ctx interface{}, data interface{}) *Loan_FindOne_Call {
	return &Loan_FindOne_Call{Call: _e.mock.On("FindOne", ctx, data)}
}

func (_c *Loan_FindOne_Call) Run(run func(ctx context.Context, data model.FindOneLoanInput)) *Loan_FindOne_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.FindOneLoanInput))
	})
	return _c
}

func (_c *Loan_FindOne_Call) Return(_a0 *databasemodel.Loan, _a1 error) *Loan_FindOne_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Loan_FindOne_Call) RunAndReturn(run func(context.Context, model.FindOneLoanInput) (*databasemodel.Loan, error)) *Loan_FindOne_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, data
func (_m *Loan) Update(ctx context.Context, data model.UpdateLoanInput) error {
	ret := _m.Called(ctx, data)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.UpdateLoanInput) error); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Loan_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type Loan_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - data model.UpdateLoanInput
func (_e *Loan_Expecter) Update(ctx interface{}, data interface{}) *Loan_Update_Call {
	return &Loan_Update_Call{Call: _e.mock.On("Update", ctx, data)}
}

func (_c *Loan_Update_Call) Run(run func(ctx context.Context, data model.UpdateLoanInput)) *Loan_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.UpdateLoanInput))
	})
	return _c
}

func (_c *Loan_Update_Call) Return(_a0 error) *Loan_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Loan_Update_Call) RunAndReturn(run func(context.Context, model.UpdateLoanInput) error) *Loan_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewLoan creates a new instance of Loan. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLoan(t interface {
	mock.TestingT
	Cleanup(func())
}) *Loan {
	mock := &Loan{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
