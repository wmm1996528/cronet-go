package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/sagernet/cronet-go"
)

func main() {
	client := &http.Client{
		Transport: &cronet.RoundTripper{},
	}
	resp, err := client.Get(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}
