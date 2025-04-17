package api

import (
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestGetUserActivity_Non200Status(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Not Found"}`))
	}))
	defer mockServer.Close()

	client := &http.Client{
		Transport: &http.Transport{
			Proxy:                 http.ProxyFromEnvironment,
			DialContext:           (&net.Dialer{}).DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
	originalClient := http.DefaultClient
	http.DefaultClient = client
	defer func() { http.DefaultClient = originalClient }()

	_, err := GetUserActivity("testuser", "dummy_token")
	if err == nil {
		t.Error("Expected error for non-200 status, got nil")
	}
	if !strings.Contains(err.Error(), "API request failed with status: 404 Not Found") {
		t.Errorf("Expected error message containing 'API request failed with status: 404 Not Found', got %v", err)
	}
}

func TestGetUserActivity_InvalidJSON(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{invalid json`))
	}))
	defer mockServer.Close()

	client := &http.Client{
		Transport: &http.Transport{
			Proxy:                 http.ProxyFromEnvironment,
			DialContext:           (&net.Dialer{}).DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
	originalClient := http.DefaultClient
	http.DefaultClient = client
	defer func() { http.DefaultClient = originalClient }()

	_, err := GetUserActivity("testuser", "dummy_token")
	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}
	if !strings.Contains(err.Error(), "failed to decode response") {
		t.Errorf("Expected error message containing 'failed to decode response', got %v", err)
	}
}

func TestGetUserActivity_RequestFailure(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	mockServer.Close()

	client := &http.Client{
		Transport: &http.Transport{
			Proxy:                 http.ProxyFromEnvironment,
			DialContext:           (&net.Dialer{}).DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}
	originalClient := http.DefaultClient
	http.DefaultClient = client
	defer func() { http.DefaultClient = originalClient }()

	_, err := GetUserActivity("testuser", "dummy_token")
	if err == nil {
		t.Error("Expected error for request failure, got nil")
	}
	if !strings.Contains(err.Error(), "failed to send request") {
		t.Errorf("Expected error message containing 'failed to send request', got %v", err)
	}
}
