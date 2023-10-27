package main

import (
	"flag"
	"github.com/weblifeio/cronet-go"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	var url string
	flag.StringVar(&url, "url", "", "HTTP/2 URL to open")
	var proxy string
	flag.StringVar(&proxy, "proxy", "", "proxy server to use")
	flag.Parse()
	if len(url) == 0 {
		log.Fatal("URL argument is not provided")
	}

	// Allocate resources
	engineParams := cronet.NewEngineParams()
	engineParams.SetUserAgent("Go-http-client/2")
	engineParams.SetProxyServer(proxy)

	// Start Cronet engine
	engine := cronet.NewEngine()
	engine.StartWithParams(engineParams)
	engine.StartNetLogToFile("netlog.json", true)

	defer func() {
		// Shutdown Cronet engine
		engine.StopNetLog()
		engine.Shutdown()

		// Clearing up resources used
		engine.Destroy()
		engineParams.Destroy()
	}()

	streamEngine := engine.StreamEngine()

	// Open HTTP2 URL
	headers := make(map[string]string)
	conn := streamEngine.CreateConn(true, true)
	err := conn.Start(http.MethodGet, url, headers, 0, true)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(os.Stdout, conn)
	if err != nil {
		log.Fatal(err)
	}
}
