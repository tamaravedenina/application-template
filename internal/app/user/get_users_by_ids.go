package module

import (
	"context"
	"errors"

	pb "application-template/vendor.pb/user/api"
)

// GetBooksByIDs ...
func (m *UserModule) GetUsersByIDs(ctx context.Context, req *pb.GetUsersByIDsRequest) (*pb.GetUsersByIDsResponse, error) {
	if len(req.UserIds) == 0 {
		return nil, errors.New("user ids is required")
	}

	users, err := m.service.GetUsersByIDs(ctx, req.UserIds)
	if err != nil {
		return nil, err
	}

	usersResponse := make([]*pb.GetUsersByIDsResponse_User, 0)
	for _, user := range users {
		usersResponse = append(usersResponse, &pb.GetUsersByIDsResponse_User{
			Id:   user.ID,
			Name: user.Name,
		})
	}

	return &pb.GetUsersByIDsResponse{Users: usersResponse}, nil
}
