package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/weblifeio/cronet-go"
)

// URL wraps the net/url.URL type to enable unmarshalling from text
type URL struct {
	url.URL
}

func (u *URL) UnmarshalText(text []byte) error {
	parsed, err := url.Parse(string(text))
	if err != nil {
		return err
	}
	u.URL = *parsed
	return nil
}

func (u *URL) MarshalText() ([]byte, error) {
	return []byte(u.String()), nil
}

func ConfigureClientCertificate(e *cronet.Engine, certPath string, keyPath string, hostPort []string) {
	if certPath == "" || keyPath == "" {
		return
	}
	clientCertData, err := os.ReadFile(certPath)
	if err != nil {
		log.Fatal(err)
	}
	privateKeyData, err := os.ReadFile(keyPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, hostPortPair := range hostPort {
		e.SetClientCertificate(hostPortPair, clientCertData, privateKeyData)
	}
}

func main() {
	var urlArg URL
	flag.TextVar(&urlArg, "url", &URL{}, "URL to open")
	var proxyArg URL
	flag.TextVar(&proxyArg, "proxy", &URL{}, "proxy server to use")
	var certPath string
	flag.StringVar(&certPath, "cert", "", "path to certificate file")
	var keyPath string
	flag.StringVar(&keyPath, "key", "", "path to key file")
	flag.Parse()
	if len(urlArg.String()) == 0 {
		log.Fatal("URL argument is not provided")
	}

	engineParams := cronet.NewEngineParams()
	engineParams.SetUserAgent("Go-http-client/2")
	engineParams.SetProxyServer(proxyArg.String())

	t := cronet.NewCronetTransport(engineParams, true)
	ConfigureClientCertificate(&t.Engine, certPath, keyPath, []string{urlArg.Host, proxyArg.Host})

	client := &http.Client{
		Transport: t,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Get(urlArg.String())
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}
