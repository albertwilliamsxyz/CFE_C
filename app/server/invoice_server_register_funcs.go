package server

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	api_pb "CFE_Clone/api/gen/go/api/protos"
)

// RegisterWithServer implements grapiserver.Server.RegisterWithServer.
func (s *invoiceServiceServerImpl) RegisterWithServer(grpcSvr *grpc.Server) {
	api_pb.RegisterInvoiceServiceServer(grpcSvr, s)
}

// RegisterWithHandler implements grapiserver.Server.RegisterWithHandler.
func (s *invoiceServiceServerImpl) RegisterWithHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return api_pb.RegisterInvoiceServiceHandler(ctx, mux, conn)
}
