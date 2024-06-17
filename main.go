package main

import (
	"fmt"
	"io"
	"log"
)

type MySlowReader struct {
	contents string
	pos      int
}

func (m MySlowReader) Read(p []byte) (n int, err error) {
	if m.pos+1 <= len(m.contents) {
		n := copy(p, m.contents[m.pos:m.pos+1])
		m.pos++
		return n, nil
	}
	return 0, io.EOF
}

func main() {

	mySlowerReaderInstance := MySlowReader{
		contents: "Hello World",
	}

	out, err := io.ReadAll(mySlowerReaderInstance)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("output: %s\n", out)
}
