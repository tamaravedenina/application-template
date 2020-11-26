package user

import (
	"application-template/internal/pkg/config"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"context"
)

// Client ...
type Client struct {
	conn UserClient
}

// NewClient ...
func NewClient() ClientInterface {
	conn, err := grpc.Dial(config.GetCfg().ApiClients["user"].Target, grpc.WithInsecure())
	if err != nil {
		logrus.Fatal(err)
	}
	return Client{
		conn: NewUserClient(conn),
	}
}

// ClientInterface ...
type ClientInterface interface {
	GetUsersByIDs(ctx context.Context, ids []int64) (map[int64]*GetUsersByIDsResponse_User, error)
}

// GetUsersByIDs ...
func (c Client) GetUsersByIDs(ctx context.Context, ids []int64) (map[int64]*GetUsersByIDsResponse_User, error) {
	result := make(map[int64]*GetUsersByIDsResponse_User)
	userResponse, err := c.conn.GetUsersByIDs(ctx, &GetUsersByIDsRequest{UserIds: ids})
	if err != nil {
		return result, err
	}

	if userResponse != nil {
		for _, user := range userResponse.Users {
			result[user.Id] = user
		}
	}

	return result, nil
}
