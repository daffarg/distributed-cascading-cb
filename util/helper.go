package util

import (
	"fmt"
	"github.com/btcsuite/btcd/btcutil/base58"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func GetEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

func GetIntEnv(key string, fallback int) int {
	value := GetEnv(key, strconv.Itoa(fallback))
	valueAsInt, _ := strconv.Atoi(value)

	return valueAsInt
}

func GetGeneralURLFormat(urlStr string) (string, error) {
	parsedUrl, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}

	// Get the host (netloc) and remove "www." if present
	host := strings.Replace(parsedUrl.Host, "www.", "", 1)

	// Get the path and ensure it does not end with a "/"
	path := strings.TrimSuffix(parsedUrl.Path, "/")

	// Combine host and path for the general form
	return host + path, nil
}

func FormEndpointName(url, method string) string {
	return fmt.Sprintf("%s:%s", method, url)
}

func FormEndpointStatusKey(endpointName string) string {
	return fmt.Sprintf("%s%s", StatusKeyPrefix, endpointName)
}

func FormRequiringEndpointsKey(endpointName string) string {
	return fmt.Sprintf("%s%s", RequiringsEndpointKeyPrefix, endpointName)
}

func EncodeTopic(topic string) string {
	return base58.Encode([]byte(topic))
}

func GetEndpointFromRequiringsKey(key string) string {
	colonIndex := strings.Index(key, ":")
	if colonIndex == -1 {
		return ""
	}

	return key[colonIndex+1:]
}
