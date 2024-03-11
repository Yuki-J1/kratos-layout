package server

import (
	swaggerUI "github.com/tx7do/kratos-swagger-ui"
	v1 "hellword/api/helloworld/v1"
	"hellword/internal/conf"
	"hellword/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)

	swaggerUI.RegisterSwaggerUIServerWithOption(
		srv,
		swaggerUI.WithTitle("helloworld v1"),
		swaggerUI.WithMemoryData(v1.OpenApiData, "yaml"),
		swaggerUI.WithBasePath("/docs/"),
	)

	v1.RegisterGreeterHTTPServer(srv, greeter)
	return srv
}
