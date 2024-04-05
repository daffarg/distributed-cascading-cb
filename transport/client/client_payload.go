package client

import (
	"context"
	"github.com/daffarg/distributed-cascading-cb/protobuf"
	"github.com/daffarg/distributed-cascading-cb/service"
)

func encodeGeneralRequestReq(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*service.GeneralRequestReq)
	return &protobuf.GeneralRequestInput{
		Method:            req.Method,
		Url:               req.URL,
		Header:            req.Header,
		Body:              req.Body,
		RequiringEndpoint: req.RequiringEndpoint,
		RequiringMethod:   req.RequiringMethod,
	}, nil
}

func decodeResponse(_ context.Context, grpcRes interface{}) (interface{}, error) {
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
