package service

import (
	"application-template/internal/pkg/gapi/user"
	"application-template/internal/pkg/helper"
	"context"
	"errors"
	"fmt"

	"application-template/internal/app/library/datastruct"
)

// GetBooksByIDs ...
func (s *LibraryService) GetBooksByIDs(ctx context.Context, ids []int64) ([]*datastruct.Book, error) {
	return s.repository.GetBooksByIDs(ctx, ids)
}

// SaveBook ...
func (s *LibraryService) SaveBook(ctx context.Context, book datastruct.Book) (*datastruct.Book, error) {
	if book.ID > 0 {
		books, err := s.GetBooksByIDs(ctx, []int64{book.ID})
		if err != nil {
			return nil, err
		}

		if len(books) == 0 {
			return nil, errors.New(fmt.Sprintf("book #%d does not exists", book.ID))
		}

		return s.repository.UpdateBook(ctx, book)
	}

	return s.repository.InsertBook(ctx, book)
}

// GetBooksWithUserByIDs ...
func (s *LibraryService) GetBooksWithUserByIDs(ctx context.Context, ids []int64) ([]*datastruct.BookWithUser, error) {
	books, err := s.GetBooksByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}

	userIDs := make([]int64, 0)
	for _, book := range books {
		if book.UserID > 0 {
			userIDs = helper.AppendInt64Uniq(userIDs, book.UserID)
		}
	}

	userMap := make(map[int64]*user.GetUsersByIDsResponse_User)
	if len(userIDs) > 0 {
		userMap, err = s.userApi.GetUsersByIDs(ctx, userIDs)
		if err != nil {
			return nil, err
		}
	}

	booksWithUser := make([]*datastruct.BookWithUser, 0)
	for _, book := range books {
		bookWithUser := &datastruct.BookWithUser{
			ID:   book.ID,
			Name: book.Name,
		}

		if book.UserID > 0 {
			if u, ok := userMap[book.UserID]; ok == true {
				bookWithUser.UserName = u.Name
			}
		}

		booksWithUser = append(booksWithUser, bookWithUser)
	}

	return booksWithUser, nil
}
