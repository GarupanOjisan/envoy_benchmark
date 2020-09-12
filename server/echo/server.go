package echo

import (
	"context"

	"github.com/garupanojisan/envoy_benchmark/proto/echo"
)

// echo server
type Server struct{}

func (s *Server) Echo(ctx context.Context, request *echo.Echo_Request) (*echo.Echo_Response, error) {
	return &echo.Echo_Response{
		Message: request.GetMessage(),
	}, nil
}
