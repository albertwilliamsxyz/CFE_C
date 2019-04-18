package app

import (
	"CFE_Clone/app/server"
	"github.com/izumin5210/grapi/pkg/grapiserver"
)

// Run starts the grapiserver.
func Run() error {
	s := grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithGrpcAddr("tcp", ":3000"),
		grapiserver.WithGatewayAddr("tcp", ":4000"),
		grapiserver.WithServers(
			server.NewInvoiceServiceServer(),
		),
	)
	return s.Serve()
}
