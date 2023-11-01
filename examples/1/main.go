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
		Transport: cronet.NewCronetRoundTripperWithDefaultParams(),
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Get(url)
	//resp, err := client.Get(os.Args[1])
	log.Println("GOT RESPONSE FROM CRONET HTTP CLIENT: ", resp)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}
