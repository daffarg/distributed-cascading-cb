package client

import (
	"context"
	"github.com/daffarg/distributed-cascading-cb/protobuf"
	"github.com/daffarg/distributed-cascading-cb/service"
)

func encodeGeneralRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*service.GeneralRequest)
	return &protobuf.GeneralRequest{
		Method:            req.Method,
		Url:               req.URL,
		Header:            req.Header,
		Body:              req.Body,
		RequiringEndpoint: req.RequiringEndpoint,
		RequiringMethod:   req.RequiringMethod,
	}, nil
}

func encodeGetRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*service.GetRequest)
	return &protobuf.GetRequest{
		Url:               req.URL,
		Header:            req.Header,
		RequiringEndpoint: req.RequiringEndpoint,
		RequiringMethod:   req.RequiringMethod,
	}, nil
}

func encodePostRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*service.PostRequest)
	return &protobuf.PostRequest{
		Url:               req.URL,
		Header:            req.Header,
		Body:              req.Body,
		RequiringEndpoint: req.RequiringEndpoint,
		RequiringMethod:   req.RequiringMethod,
	}, nil
}

func encodePutRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*service.PutRequest)
	return &protobuf.PutRequest{
		Url:               req.URL,
		Header:            req.Header,
		Body:              req.Body,
		RequiringEndpoint: req.RequiringEndpoint,
		RequiringMethod:   req.RequiringMethod,
	}, nil
}

func encodeDeleteRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*service.DeleteRequest)
	return &protobuf.DeleteRequest{
		Url:               req.URL,
		Header:            req.Header,
		RequiringEndpoint: req.RequiringEndpoint,
		RequiringMethod:   req.RequiringMethod,
	}, nil
}

func decodeResponse(ctx context.Context, grpcRes interface{}) (interface{}, error) {
	res := grpcRes.(*protobuf.Response)
	return &service.Response{
		Status:        res.Status,
		StatusCode:    res.StatusCode,
		Proto:         res.Proto,
		ProtoMajor:    res.ProtoMajor,
		ProtoMinor:    res.ProtoMinor,
		Header:        res.Header,
		Body:          res.Body,
		ContentLength: res.ContentLength,
	}, nil
}
