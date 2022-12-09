package main

import (
	"fmt"
	"http-mocking/pkg/api"
	"http-mocking/pkg/httputil"
)

func main() {
	client := httputil.NewHttpClient() // get a REAL http client
	ip, err := api.GetIpAddress(client)
	if err != nil {
		fmt.Printf("%+v", err)
		return
	}
	fmt.Printf("ip: %s", ip)
}
