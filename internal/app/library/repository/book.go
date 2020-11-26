package repository

import (
	"application-template/internal/pkg/db"
	"context"
	"github.com/go-pg/pg"

	"application-template/internal/app/library/datastruct"
)

// GetBooksByIDs ...
func (r *LibraryRepository) GetBooksByIDs(ctx context.Context, ids []int64) ([]*datastruct.Book, error) {
	books := make([]*datastruct.Book, 0)
	err := db.GetDB(ctx).Model(&books).Where("id in (?)", pg.In(ids)).Select()
	return books, err
}

// InsertBook ...
func (r *LibraryRepository) InsertBook(ctx context.Context, book datastruct.Book) (*datastruct.Book, error) {
	_, err := db.GetDB(ctx).Model(&book).Insert()
	return &book, err
}

// UpdateBook ...
func (r *LibraryRepository) UpdateBook(ctx context.Context, book datastruct.Book) (*datastruct.Book, error) {
	_, err := db.GetDB(ctx).Model(&book).Where("id = ?", book.ID).Update()
	return &book, err
}
