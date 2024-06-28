package register

import (
	"dependency-injection/domain/user"
	"fmt"
)

type RegisterService interface {
	Register(user.User) error
}

func NewService(r Repository) RegisterService {
	return &registerService{repository: r}
}

type registerService struct {
	repository Repository
}

func (r *registerService) Register(u user.User) error {
	err := r.repository.Add(u)
	if err != nil {
		return fmt.Errorf("failed to register: %w", err)
	}
	return nil
}
