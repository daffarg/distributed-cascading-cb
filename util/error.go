package util

import "errors"

var (
	ErrFailedParsingURL       = errors.New("failed to parse URL")
	ErrKeyNotFound            = errors.New("key not found")
	ErrCircuitBreakerOpen     = errors.New("circuit breaker is open")
	ErrFailedExecuteRequest   = errors.New("failed to execute the request")
	ErrUnsupportedMessageType = errors.New("unsupported message type")
)
