package api

import (
	"bytes"
	"encoding/json"
	"http-mocking/pkg/httputil"
	"io"
	"net/http"
	"testing"
)

func TestGetIpAddress(t *testing.T) {
	myMockClient := newMockHTTPClient("20.40.60.80")
	ip, _ := GetIpAddress(myMockClient)
	if ip != "20.40.60.80" {
		t.Fatalf("Got wrong value for ip")
	}
}

func newMockHTTPClient(mockIp string) mockHTTPClient {
	return mockHTTPClient{
		ip: mockIp,
	}
}

type mockHTTPClient struct {
	httputil.Client
	ip string
}

// mocked implementation of Do()
func (m mockHTTPClient) Do(_ *http.Request) (*http.Response, error) {
	mockResponseBody := createMockResponseBody(
		ResponseBody{
			Ip: m.ip,
		})
	res := http.Response{
		StatusCode: 200,
		Body:       mockResponseBody,
	}
	return &res, nil
}

func createMockResponseBody(body ResponseBody) io.ReadCloser {
	mockResponseBytes, _ := json.Marshal(body)
	return io.NopCloser(bytes.NewReader(mockResponseBytes))
}
