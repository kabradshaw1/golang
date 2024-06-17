package main

type RequestError struct {
	HTTPCode int
	Body     string
	Err      string
}
