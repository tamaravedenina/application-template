package module

import (
	"context"
	"errors"

	pb "application-template/vendor.pb/library/api"
)

// GetBooksByIDs ...
func (m *LibraryModule) GetBooksByIDs(ctx context.Context, req *pb.GetBooksByIDsRequest) (*pb.GetBooksByIDsResponse, error) {
	if len(req.BookIds) == 0 {
		return nil, errors.New("books ids is required")
	}

	books, err := m.service.GetBooksByIDs(ctx, req.BookIds)
	if err != nil {
		return nil, err
	}

	booksResponse := make([]*pb.Book, 0)
	for _, book := range books {
		booksResponse = append(booksResponse, &pb.Book{
			Id:   book.ID,
			Name: book.Name,
		})
	}

	return &pb.GetBooksByIDsResponse{Books: booksResponse}, nil
}
