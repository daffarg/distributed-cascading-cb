package service

import (
	"bytes"
	"context"
	"github.com/daffarg/distributed-cascading-cb/util"
	"io"
	"net/http"
)

func (s *service) httpRequest(ctx context.Context, method, url string, body []byte, header map[string]string) (*Response, error) {
	client := s.httpClient
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	for k, v := range header {
		req.Header.Set(k, v)
	}

	req = req.Clone(ctx)

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if httpRes.StatusCode >= 500 {
		return nil, util.ErrFailedExecuteRequest
	}
	defer httpRes.Body.Close()

	body, err = io.ReadAll(httpRes.Body)
	if err != nil {
		return &Response{}, err
	}

	res := &Response{
		Status:        httpRes.Status,
		StatusCode:    int32(httpRes.StatusCode),
		Proto:         httpRes.Proto,
		ProtoMajor:    int32(httpRes.ProtoMajor),
		ProtoMinor:    int32(httpRes.ProtoMinor),
		Body:          body,
		ContentLength: httpRes.ContentLength,
	}

	responseHeader := make(map[string]string)
	for k := range header {
		responseHeader[k] = httpRes.Header.Get(k)
	}
	res.Header = responseHeader

	return res, nil
}
