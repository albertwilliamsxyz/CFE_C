package server

import (
	api_pb "CFE_Clone/api/gen/go/api/protos"
	"encoding/json"
	"strconv"
)

type data struct {
	Bills []bills
}

type bills struct {
	Balance float64
	Status string
	DateFrom string
	DateTo string
	Month string
	ExpirationDate string
	Anio string
}

type unmarshalledData struct {
	Status  string
	Message string
	Data    data
}

var monthLanguagesCorrespondence = map[string]api_pb.Month{
	"":           api_pb.MONTH_INVALID,
	"Enero":      api_pb.MONTH_JANUARY,
	"Febrero":    api_pb.MONTH_FEBRUARY,
	"Marzo":      api_pb.MONTH_MARCH,
	"Abril":      api_pb.MONTH_APRIL,
	"Mayo":       api_pb.MONTH_MAY,
	"Junio":      api_pb.MONTH_JUNE,
	"Julio":      api_pb.MONTH_JULY,
	"Agosto":     api_pb.MONTH_AUGUST,
	"Septiembre": api_pb.MONTH_SEPTEMBER,
	"Octubre":    api_pb.MONTH_OCTOBER,
	"Noviembre":  api_pb.MONTH_NOVEMBER,
	"Diciembre":  api_pb.MONTH_DECEMBER,
}

func stringToDate(dateString string) *api_pb.Date {
	year, _ := strconv.Atoi(dateString[:4])
	month, _ := strconv.Atoi(dateString[4:6])
	day, _ := strconv.Atoi(dateString[6:8])

	return &api_pb.Date{
		Year:  int32(year),
		Month: int32(month),
		Day:   int32(day),
	}
}

func JSONInvoicesToGRPCInvoices(JSONResponse []byte) (*api_pb.ListInvoicesResponse, error) {
	var unmarshalledData unmarshalledData
	err := json.Unmarshal(JSONResponse, &unmarshalledData);
	if err != nil {
		return nil, err
	}

	var invoices []*api_pb.Invoice
	for _, value := range unmarshalledData.Data.Bills {
		year, _ := strconv.Atoi(value.Anio)
		invoices = append(invoices, &api_pb.Invoice{
			Total:          value.Balance,
			Status:         value.Status,
			DateFrom:       stringToDate(value.DateFrom),
			DateTo:         stringToDate(value.DateTo),
			Month:          monthLanguagesCorrespondence[value.Month],
			ExpirationDate: stringToDate(value.ExpirationDate),
			Year:           int32(year),
		})
	}

	return &api_pb.ListInvoicesResponse{
		Invoices: invoices,
	}, nil
}
