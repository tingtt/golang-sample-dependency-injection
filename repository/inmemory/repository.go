package inmemory

import (
	"dependency-injection/domain/user"
	"dependency-injection/usecase/register"
	"errors"
)

func NewRepository() register.Repository {
	return &repository{}
}

type repository struct{}

func (r *repository) Add(user.User) error {
	return errors.New("unimplemented")
}
