package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/weblifeio/cronet-go"
)

func main() {
	var url string
	flag.StringVar(&url, "url", "", "URL to open")
	flag.Parse()
	if len(url) == 0 {
		log.Fatal("URL argument is not provided")
	}

	client := &http.Client{
		Transport: &cronet.RoundTripper{},
	}
	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}
