package main

import (
	"fmt"
	"github.com/weblifeio/cronet-go"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	proxy := "http://us.ipwo.net:7878"
	proxy = "https://127.0.0.1:8081"
	url := "https://httpbin.org/ip"
	// Allocate resources
	engineParams := cronet.NewEngineParams()
	engineParams.SetUserAgent("Go-http-client/2")
	engineParams.SetProxyServer(proxy)
	engineParams.SetEnableQuic(false)
	engineParams.SetEnableHTTP2(true)
	engineParams.SetEnableBrotli(true)
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
	headers := map[string]string{}
	headers["Proxy-Authorization"] = "Basic realm=dW5pMDAwMDJfY3VzdG9tX3pvbmVfREVfc2lkXzg1NjY2ODgzX3RpbWVfNTpGZGtpR0h0eTlh"
	headers["proxy-authorization"] = "Basic realm=dW5pMDAwMDJfY3VzdG9tX3pvbmVfREVfc2lkXzg1NjY2ODgzX3RpbWVfNTpGZGtpR0h0eTlh"
	headers["proxy-authenticate"] = "Basic realm=dW5pMDAwMDJfY3VzdG9tX3pvbmVfREVfc2lkXzg1NjY2ODgzX3RpbWVfNTpGZGtpR0h0eTlh"
	fmt.Println(headers)
	fmt.Println(url)
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
