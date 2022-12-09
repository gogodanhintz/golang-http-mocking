package api

import (
	"encoding/json"
	"fmt"
	"http-mocking/pkg/httputil"
)

// ResponseBody This is the structure of the response body that we expect
type ResponseBody struct {
	Ip string
}

func GetIpAddress(httpClient httputil.Client) (string, error) {
	url := "http://ip.jsontest.com/"

	// Get the response body
	responseBody, err := httputil.DoGet(httpClient, url)
	if err != nil {
		fmt.Printf("Error calling DoGet: %+v", err)
		return "?", err
	}

	// Unmarshal the response body into the structure we expect
	var data ResponseBody
	err = json.Unmarshal([]byte(responseBody), &data)
	if err != nil {
		fmt.Printf("Error unmarshalling responseBody: %+v", err)
		return "?", err
	}

	return data.Ip, nil
}
