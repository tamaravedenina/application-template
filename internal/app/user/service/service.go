package service

import (
	"application-template/internal/app/user/datastruct"
	"application-template/internal/app/user/repository"
	"context"
)

// UserService ...
type UserService struct {
	repository repository.UserRepositoryInterface
}

// NewUserService ...
func NewUserService(repository repository.UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		repository: repository,
	}
}

// BuildUserService ...
func BuildUserService() UserServiceInterface {
	libraryRepository := repository.BuildUserRepository()
	return NewUserService(libraryRepository)
}

// UserServiceInterface ...
type UserServiceInterface interface {
	GetUsersByIDs(ctx context.Context, ids []int64) ([]*datastruct.User, error)
}
