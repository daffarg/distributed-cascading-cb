package util

const (
	LogRequest                 = "request"
	LogResponse                = "response"
	LogError                   = "error"
	LogMessage                 = "message"
	LogTimestamp               = "timestamp"
	LogPath                    = "path"
	LogTopic                   = "topic"
	LogCircuitBreakerEndpoint  = "circuit_breaker_endpoint"
	LogCircuitBreakerOldStatus = "circuit_breaker_old_status"
	LogCircuitBreakerNewStatus = "circuit_breaker_new_status"
	LogAlternativeEndpoint     = "alternative_endpoint"
)

const (
	RequiredEndpointsKey = "required_endpoints"
)

const (
	Get    = "GET"
	Post   = "POST"
	Put    = "PUT"
	Delete = "DELETE"
)
