package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginRequest struct {
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func doLoginRequest(requestURL string, password string) (string, error) {
	loginRequest := LoginRequest{
		Password: password,
	}

	body, err := json.Marshal((loginRequest))
	if err != nil {
		return "", fmt.Errorf("json.Marshal error: %s", err)
	}

	http.Post(requestURL, "application/json", bytes.NewBuffer(body))
	return "", nil
}
