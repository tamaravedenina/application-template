package module

import (
	"application-template/internal/app/library/datastruct"
	"context"
	"errors"

	pb "application-template/vendor.pb/library/api"
)

// SaveBook ...
func (m *LibraryModule) SaveBook(ctx context.Context, req *pb.SaveBooksRequest) (*pb.SaveBooksResponse, error) {
	if len(req.Name) == 0 {
		return nil, errors.New("book name is required")
	}

	book, err := m.service.SaveBook(ctx, datastruct.Book{ID: req.Id, Name: req.Name})
	if err != nil {
		return nil, err
	}
	return &pb.SaveBooksResponse{Book: &pb.Book{Id: book.ID, Name: book.Name}}, nil
}
