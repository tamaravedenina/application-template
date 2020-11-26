package module

import (
	"context"
	"errors"

	pb "application-template/vendor.pb/library/api"
)

// GetBooksByIDs ...
func (m *LibraryModule) GetBooksWithUserByIDs(ctx context.Context, req *pb.GetBooksWithUserByIDsRequest) (*pb.GetBooksWithUserByIDsResponse, error) {
	if len(req.BookIds) == 0 {
		return nil, errors.New("books ids is required")
	}

	booksWithUser, err := m.service.GetBooksWithUserByIDs(ctx, req.BookIds)
	if err != nil {
		return nil, err
	}

	booksResponse := make([]*pb.GetBooksWithUserByIDsResponse_BookWithUser, 0)
	for _, bookWithUser := range booksWithUser {
		booksResponse = append(booksResponse, &pb.GetBooksWithUserByIDsResponse_BookWithUser{
			Id:       bookWithUser.ID,
			Name:     bookWithUser.Name,
			UserName: bookWithUser.UserName,
		})
	}

	return &pb.GetBooksWithUserByIDsResponse{Books: booksResponse}, nil
}
