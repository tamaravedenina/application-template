package repository

import (
	"context"

	"application-template/internal/app/library/datastruct"
)

// LibraryRepository ...
type LibraryRepository struct{}

// NewLibraryRepository ...
func NewLibraryRepository() LibraryRepositoryInterface {
	return &LibraryRepository{}
}

// BuildLibraryRepository ...
func BuildLibraryRepository() LibraryRepositoryInterface {
	return NewLibraryRepository()
}

// LibraryRepositoryInterface ...
type LibraryRepositoryInterface interface {
	GetBooksByIDs(ctx context.Context, ids []int64) ([]*datastruct.Book, error)
	InsertBook(ctx context.Context, book datastruct.Book) (*datastruct.Book, error)
	UpdateBook(ctx context.Context, book datastruct.Book) (*datastruct.Book, error)
}
