package server

import (
	"fmt"
	"github.com/go-openapi/spec"
	"github.com/pkg/errors"
	"github.com/utrack/clay/v2/transport"
)

func generateSwagger(descSlice []transport.ServiceDesc) ([]byte, error) {
	var swagger spec.Swagger
	for _, desc := range descSlice {
		var swg spec.Swagger
		err := swg.UnmarshalJSON(desc.SwaggerDef())
		if err != nil {
			return []byte{}, errors.New(fmt.Sprintf("couldn't UnmarshalJSON swagger: %s", err.Error()))
		}

		if swagger.Paths == nil {
			swagger = swg
		} else {
			for key, path := range swg.Paths.Paths {
				swagger.Paths.Paths[key] = path
			}
			for key, def := range swg.Definitions {
				swagger.Definitions[key] = def
			}
			for key, param := range swg.Parameters {
				swagger.Parameters[key] = param
			}
			for key, res := range swg.Responses {
				swagger.Responses[key] = res
			}
			swagger.Tags = append(swagger.Tags, swg.Tags...)
			swagger.Info.Title += " & " + swg.Info.Title
		}
	}

	return swagger.MarshalJSON()
}
