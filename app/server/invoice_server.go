package server

import (
	"context"
	"github.com/golang/glog"

	api_pb "CFE_Clone/api/gen/go/api/protos"
	"github.com/izumin5210/grapi/pkg/grapiserver"
)

// NewInvoiceServiceServer creates a new InvoiceServiceServer instance.
func NewInvoiceServiceServer() interface {
	api_pb.InvoiceServiceServer
	grapiserver.Server
} {
	return &invoiceServiceServerImpl{
		AccessToken: "Bearer fee0801a-b449-47a7-9b54-11b3cdbbdc6e",
	}
}

type invoiceServiceServerImpl struct {
	AccessToken string
}

func (s *invoiceServiceServerImpl) ListInvoices(ctx context.Context, req *api_pb.ListInvoicesRequest) (*api_pb.ListInvoicesResponse, error) {
	glog.Info("You requested: ", req.ServiceNumber)
	response, err := ListInvoices(req, s.AccessToken)
	if err != nil {
		return nil, err
	}
	return response, nil
}