package transport

import (
	"context"

	"github.com/daffarg/distributed-cascading-cb/protobuf"
	"github.com/daffarg/distributed-cascading-cb/service"
)

func decodeGeneralRequestReq(_ context.Context, r interface{}) (interface{}, error) {
	pbReq := r.(*protobuf.GeneralRequestInput)

	return &service.GeneralRequestReq{
		Method:            pbReq.Method,
		URL:               pbReq.Url,
		Header:            pbReq.Header,
		Body:              pbReq.Body,
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
