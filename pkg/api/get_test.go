package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
)

type MockClient struct {
	ResponseOutput *http.Response
}

func (m *MockClient) Get(url string) (resp *http.Response, err error) {
	return m.ResponseOutput, nil
}
func TestDoGetRequest(t *testing.T) {
	words := WordsPage{
		Page: Page{"words"},
		Words: Words{
			Input: "ab", Words: []string{"a", "b"},
		},
	}

	wordsBytes, err := json.Marshal(words)
	if err != nil {
		t.Fatalf("Marshal error: %s", err)
	}

	apiInstance := api{
		Options: Options{},
		Client: &MockClient{
			ResponseOutput: &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader(wordsBytes)),
			},
		},
	}
	response, err := apiInstance.DoGetRequest("http://locahost/words")
	if err != nil {
		t.Errorf("DoGetRequest error: %s", err)
	}
	if response == nil {
		t.Fatalf("Response is nil")
	}
	if response.GetResponse() != strings.Join([]string{"a", "b"}, ", ") {
		t.Errorf("Got wrong output: %s", response.GetResponse())
	}
}
