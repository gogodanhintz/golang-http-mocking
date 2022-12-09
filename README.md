# golang-http-mocking

This repo shows a simple example of how to create a mock for a http client.

### main.go

Running this file will call `GetIpAddress()` in the api package,
using a real http client.  So the response will be your actual IP address.

### httputil/client.go

This file defines the generic interface for a Client and which methods
we want to use and mock.  Currently, we only care about the Do() method
on http.

We then create a helper function, `DoGet()`, which takes a Client
(could be real, could be a mock) and a URL, and calls `Do()`.
This could be expanded to take request headers as well.

### api/api.go

This file has the implementation of `GetIpAddress()` that is called from
`main.go` and our unit test.

It simply calls the `DoGet()` helper function to get
a response body, then marshals that into a `ResponseBody` struct
that we are expecting.

### api/api_test.go

This file contains our unit test.

We first create a `newMockHTTPClient` and pass in the ip that we want to mock in the response.
Then, pass that client into `GetIpAddress()` and verify the response.
