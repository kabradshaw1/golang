package main

import "net/http"

type MyJWTTransport struct {
	transport http.RoundTripper
	token     string
}

func (m MyJWTTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.transport.RoundTrip(req)
}
