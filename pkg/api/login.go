package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LoginRequest struct {
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func doLoginRequest(client http.Client, requestURL string, password string) (string, error) {
	loginRequest := LoginRequest{
		Password: password,
	}

	body, err := json.Marshal((loginRequest))
	if err != nil {
		return "", fmt.Errorf("json.Marshal error: %s", err)
	}

	response, err := client.Post(requestURL, "application/json", bytes.NewBuffer(body))

	if err != nil {
		return "", fmt.Errorf("post error: %s", err)
	}

	defer response.Body.Close()

	resBody, err := io.ReadAll(response.Body)

	if err != nil {
		return "", fmt.Errorf("ReadAll error: %s", err)
	}

	if response.StatusCode != 200 {
		return "", fmt.Errorf("invalid output (HTTP Code %d): %s", response.StatusCode, string(body))
	}

	if !json.Valid(resBody) {
		return "", RequestError{
			Err:      "response is not a json",
			HTTPCode: response.StatusCode,
			Body:     string(resBody),
		}
	}

	var loginResponse LoginResponse

	err = json.Unmarshal(resBody, &loginResponse)
	if err != nil {
		return "", RequestError{
			Err:      fmt.Sprintf("Page unmarshal error: %s", err),
			HTTPCode: response.StatusCode,
			Body:     string(resBody),
		}
	}

	if loginResponse.Token == "" {
		return "", RequestError{
			HTTPCode: response.StatusCode,

			Body: string(resBody),
			Err:  "empty token"}
	}

	return loginResponse.Token, nil
}
