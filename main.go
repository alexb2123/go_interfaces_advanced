package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type logWriter struct{}

func main() {

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	lw := logWriter{}

	_, err = io.Copy(lw, file)
	if err != nil {
		log.Fatal(err)
	}

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
