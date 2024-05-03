package transport

import (
	"context"

	"github.com/daffarg/distributed-cascading-cb/protobuf"
	"github.com/daffarg/distributed-cascading-cb/service"
)

func decodeGeneralRequest(_ context.Context, r interface{}) (interface{}, error) {
	pbReq := r.(*protobuf.GeneralRequest)

	return &service.GeneralRequest{
		Method:            pbReq.Method,
		URL:               pbReq.Url,
		Header:            pbReq.Header,
		Body:              pbReq.Body,
		RequiringEndpoint: pbReq.RequiringEndpoint,
		RequiringMethod:   pbReq.RequiringMethod,
	}, nil
}

func decodeGetRequest(_ context.Context, r interface{}) (interface{}, error) {
	pbReq := r.(*protobuf.GetRequest)

	return &service.GetRequest{
		URL:               pbReq.Url,
		Header:            pbReq.Header,
		RequiringEndpoint: pbReq.RequiringEndpoint,
		RequiringMethod:   pbReq.RequiringMethod,
	}, nil
}

func decodePostRequest(_ context.Context, r interface{}) (interface{}, error) {
	pbReq := r.(*protobuf.PostRequest)

	return &service.PostRequest{
		URL:               pbReq.Url,
		Header:            pbReq.Header,
		Body:              pbReq.Body,
		RequiringEndpoint: pbReq.RequiringEndpoint,
		RequiringMethod:   pbReq.RequiringMethod,
	}, nil
}

func decodePutRequest(_ context.Context, r interface{}) (interface{}, error) {
	pbReq := r.(*protobuf.PutRequest)

	return &service.PutRequest{
		URL:               pbReq.Url,
		Header:            pbReq.Header,
		Body:              pbReq.Body,
		RequiringEndpoint: pbReq.RequiringEndpoint,
		RequiringMethod:   pbReq.RequiringMethod,
	}, nil
}

func decodeDeleteRequest(_ context.Context, r interface{}) (interface{}, error) {
	pbReq := r.(*protobuf.PutRequest)

	return &service.DeleteRequest{
		URL:               pbReq.Url,
		Header:            pbReq.Header,
		RequiringEndpoint: pbReq.RequiringEndpoint,
		RequiringMethod:   pbReq.RequiringMethod,
	}, nil
}

func encodeResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(*service.Response)

	return &protobuf.Response{
		Status:        res.Status,
		StatusCode:    res.StatusCode,
		Body:          res.Body,
		Header:        res.Header,
		Proto:         res.Proto,
		ProtoMajor:    res.ProtoMajor,
		ProtoMinor:    res.ProtoMinor,
		ContentLength: res.ContentLength,
	}, nil
}
