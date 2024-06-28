package register

import (
	"dependency-injection/domain/user"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Add(u user.User) error {
	args := m.Called(u)
	return args.Error(0)
}

type args struct {
	u user.User
}

func Test_registerService_Register_may_success(t *testing.T) {
	t.Parallel()

	repository := new(MockRepository)
	repository.On("Add", mock.Anything).Return(nil)

	tests := []struct {
		repository Repository
		args       args
	}{
		{
			repository: repository,
			args: args{
				u: user.User{},
			},
		},
	}
	for i, tt := range tests {
		t.Run(
			fmt.Sprintf("success %d", i+1),
			func(t *testing.T) {
				r := &registerService{tt.repository}
				err := r.Register(tt.args.u)
				if err != nil {
					t.Errorf("registerService.Register() error = %vv", err)
				}
			})
	}
}

func Test_registerService_Register_may_timeout(t *testing.T) {
	t.Parallel()

	repository := new(MockRepository)
	repository.On("Add", mock.Anything).Return(ErrTimeout)

	tests := []struct {
		repository Repository
		args       args
	}{
		{
			repository: repository,
			args: args{
				u: user.User{},
			},
		},
	}
	for i, tt := range tests {
		t.Run(
			fmt.Sprintf("fail %d", i+1),
			func(t *testing.T) {
				r := &registerService{tt.repository}
				err := r.Register(tt.args.u)
				if err != nil && !errors.Is(err, ErrTimeout) {
					t.Errorf("registerService.Register() error = %v, wantErr %v", err, ErrTimeout)
				}
			})
	}
}
