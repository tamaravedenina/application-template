package library

import (
	"context"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"application-template/internal/pkg/config"
)

// Client book client
type Client struct {
	conn LibraryClient
}

// NewClient returns book client
func NewClient() Client {
	conn, err := grpc.Dial(config.GetCfg().ApiClients["library"].Target, grpc.WithInsecure())
	if err != nil {
		logrus.Fatal(err)
	}
	return Client{
		conn: NewLibraryClient(conn),
	}
}

// IBookClient book client interface
type IBookClient interface {
	GetUsersByIDs(ctx context.Context, IDs []int64) (map[int64]*GetBooksByUserIDResponse_Book, error)
}

// GetUsersByIDs returns user's books map
func (c Client) GetUsersByIDs(ctx context.Context, IDs []int64) (map[int64]*GetBooksByUserIDResponse_Book, error) {
	result := make(map[int64]*GetBooksByUserIDResponse_Book)
	bookResponse, err := c.conn.GetBooksByUserID(ctx, &GetBooksByUserIDRequest{UserIds: IDs})
	if err != nil {
		return result, err
	}

	if bookResponse != nil {
		for _, user := range bookResponse.Users {
			result[user.Id] = user
		}
	}

	return result, nil
}
