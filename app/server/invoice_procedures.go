package server

import (
	api_pb "CFE_Clone/api/gen/go/api/protos"
	"bytes"
	"fmt"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
	"net/url"
)

func ListInvoices(request *api_pb.ListInvoicesRequest, accessToken string) (*api_pb.ListInvoicesResponse, error) {
	data, err := RetrieveJSONInvoicesByService(request.ServiceNumber, accessToken)
	if err != nil {
		return nil, fmt.Errorf("RetrieveJSONInvoicesByService: %v", err)
	}

	response, err := JSONInvoicesToGRPCInvoices(data)
	if err != nil {
		return nil, fmt.Errorf("JSONInvoicesToGRPCInvoices: %v", err)
	}

	return response, nil
}

func RetrieveJSONInvoicesByService(serviceNumber string, accessToken string) ([]byte, error) {
	serviceURL := fmt.Sprintf("%v://%v:%v/%v", CFEProtocol, CFEHostName, CFEPort, Path[PathGetInvoices])

	requestBody := url.Values{}
	requestBody.Add("idusuario", CFEOperationalUserId)
	requestBody.Add("numerodeservicio", serviceNumber)

	request, err := http.NewRequest("POST", serviceURL, bytes.NewBufferString(requestBody.Encode()))

	if err != nil {
		return nil, fmt.Errorf("Request generation: %v", err)
	}

	request.Header.Set("Authorization", accessToken)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("Client: %v", err)
	}

	data, _ := ioutil.ReadAll(response.Body)
	return data, nil
}

func GetAccessToken() string {
	serviceURL := fmt.Sprintf("%v://%v:%v/%v", CFEProtocol, CFEHostName, CFEPort, Path[PathGetToken])

	requestBody := url.Values{}
	requestBody.Set("idusuario", CFEOperationalUserId)
	requestBody.Set("grant_type", CFEHTTPGrantType)
	requestBody.Set("username", CFEOperationalUserId)
	requestBody.Set("password", CFEOperationalUserPassword)
	requestBody.Set("scope", CFEHTTPScope)

	request, err := http.NewRequest("POST", serviceURL, bytes.NewBufferString(requestBody.Encode()))

	if err != nil {
		glog.Infof("Request generation: %v\n", err)
	}

	request.Header.Add("Authorization", "Basic Y2ZlLWNvbnRpZ28tY2xpZW50OmNmMy1jb250aWdvLXMzY3IzdA==")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		glog.Infof("Client: %v\n", err)
	}

	data, _ := ioutil.ReadAll(response.Body)
	return string(data)
}