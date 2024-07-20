package api

import (
	"net/http"
	"testing"
)

type MockClient struct {
	ResponseOutput *http.Response
}

func (m *MockClient) Get(url string) (resp *http.Response, err error) {
	return m.ResponseOutput, nil
}
func TestDoGetRequest(t *testing.T) {
	apiInstance := api{
		Options: Options{},
		Client: &MockClient{
			ResponseOutput: &http.Response{
				StatusCode: 200,
			},
		},
	}
	apiInstance.DoGetRequest("http://locahost/words")
}
