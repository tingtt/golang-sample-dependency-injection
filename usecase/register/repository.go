package register

import (
	"dependency-injection/domain/user"
	"errors"
)

var (
	ErrUsernameAlreadyUsed = errors.New("username already used")
	ErrTimeout             = errors.New("timeout")
)

type Repository interface {
	Add(user.User) error
}
