package server

import (
	"context"
	"testing"

	"CFE_Clone/api"
)

func Test_InvoiceServiceServer_ListInvoices(t *testing.T) {
	svr := NewInvoiceServiceServer()

	ctx := context.Background()
	req := &api_pb.ListInvoicesRequest{}

	resp, err := svr.ListInvoices(ctx, req)

	if err != nil {
		t.Errorf("returned an error %v", err)
	}

	if resp == nil {
		t.Error("response should not nil")
	}
}