package module

import (
	"application-template/internal/app/library/service"
	"github.com/utrack/clay/v2/transport"

	pb "application-template/vendor.pb/library/api"
)

// LibraryModule ...
type LibraryModule struct {
	service service.LibraryServiceInterface
}

// NewLibraryModule ...
func NewLibraryModule(service service.LibraryServiceInterface) *LibraryModule {
	return &LibraryModule{
		service: service,
	}
}

// BuildLibraryModule ...
func BuildLibraryModule() *LibraryModule {
	libraryService := service.BuildLibraryService()
	return NewLibraryModule(libraryService)
}

// GetDescription ...
func (m *LibraryModule) GetDescription() transport.ServiceDesc {
	return pb.NewLibraryServiceDesc(m)
}
