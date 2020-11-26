package module

import (
	"application-template/internal/app/user/service"
	pb "application-template/vendor.pb/user/api"
	"github.com/utrack/clay/v2/transport"
)

// UserModule ...
type UserModule struct {
	service service.UserServiceInterface
}

// NewUserModule ...
func NewUserModule(service service.UserServiceInterface) *UserModule {
	return &UserModule{
		service: service,
	}
}

// BuildUserModule ...
func BuildUserModule() *UserModule {
	userService := service.BuildUserService()
	return NewUserModule(userService)
}

// GetDescription ...
func (m *UserModule) GetDescription() transport.ServiceDesc {
	return pb.NewUserServiceDesc(m)
}
