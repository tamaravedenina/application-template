package service

import (
	"application-template/internal/app/library/datastruct"
	"application-template/internal/app/library/repository"
	"application-template/internal/pkg/gapi/user"
	"context"
)

// LibraryService ...
type LibraryService struct {
	repository repository.LibraryRepositoryInterface
	userApi    user.ClientInterface
}

// NewLibraryService ...
func NewLibraryService(repository repository.LibraryRepositoryInterface, userApi user.ClientInterface) LibraryServiceInterface {
	return &LibraryService{
		repository: repository,
		userApi:    userApi,
	}
}

// BuildLibraryService ...
func BuildLibraryService() LibraryServiceInterface {
	libraryRepository := repository.BuildLibraryRepository()
	userApi := user.NewClient()
	return NewLibraryService(libraryRepository, userApi)
}

// LibraryServiceInterface ...
type LibraryServiceInterface interface {
	GetBooksByIDs(ctx context.Context, ids []int64) ([]*datastruct.Book, error)
	GetBooksWithUserByIDs(ctx context.Context, ids []int64) ([]*datastruct.BookWithUser, error)
	SaveBook(ctx context.Context, book datastruct.Book) (*datastruct.Book, error)
}
