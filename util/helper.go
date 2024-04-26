package util

import (
	"bufio"
	"fmt"
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
	return fmt.Sprintf("status:%s", endpointName)
}

func FormRequiringEndpointsKey(endpointName string) string {
	return fmt.Sprintf("requirings:%s", endpointName)
}

func ReadPropertiesFile(path string) (map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) < 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		config[key] = value
	}

	return config, nil
}
