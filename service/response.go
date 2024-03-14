package service

import (
	"io"
	"net/http"
)

type Response struct {
	Status        string            `json:"status"`
	StatusCode    int32             `json:"status_code"`
	Proto         string            `json:"proto"`
	ProtoMajor    int32             `json:"proto_major"`
	ProtoMinor    int32             `json:"proto_minor"`
	Header        map[string]string `json:"header"`
	Body          []byte            `json:"body"`
	ContentLength int64             `json:"content_length"`
}

func (s *service) convertToResponse(res *http.Response) (Response, error) {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Response{}, err
	}

	return Response{
		Status:        res.Status,
		StatusCode:    int32(res.StatusCode),
		Proto:         res.Proto,
		ProtoMajor:    int32(res.ProtoMajor),
		ProtoMinor:    int32(res.ProtoMinor),
		Header:        s.convertResponseHeader(res.Header),
		Body:          body,
		ContentLength: res.ContentLength,
	}, nil
}

func (s *service) convertResponseHeader(header http.Header) map[string]string {
	responseHeader := make(map[string]string)
	for k := range header {
		responseHeader[k] = header.Get(k)
	}
	return responseHeader
}
