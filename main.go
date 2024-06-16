package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

// {"page":"words","input":"","words":[]}

type Words struct {
	Page  string   `json:"page"`
	Input string   `json:"input"`
	Word  []string `json:"words"`
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Printf("Usage: ./http-get <url>\n")
		os.Exit(1)
	}

	if _, err := url.ParseRequestURI(args[1]); err != nil {
		fmt.Printf("Invalid URL %s\n", err)
		os.Exit(1)
	}

	response, err := http.Get(args[1])
	if err != nil {
		log.Fatal((err))
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("HTTP Status Code: %d\nBody: %s", response.StatusCode, body)
}
