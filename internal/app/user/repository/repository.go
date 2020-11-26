package repository

import (
	"context"

	"application-template/internal/app/user/datastruct"
)

// UserRepository ...
type UserRepository struct{}

// NewLibraryRepository ...
func NewUserRepository() UserRepositoryInterface {
	return &UserRepository{}
}

// BuildLibraryRepository ...
func BuildUserRepository() UserRepositoryInterface {
	return NewUserRepository()
}

// UserRepositoryInterface ...
type UserRepositoryInterface interface {
	GetUsersByIDs(ctx context.Context, ids []int64) ([]*datastruct.User, error)
}
