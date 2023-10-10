package main

import (
	"github.com/weblifeio/cronet-go"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// Allocate resources
	engineParams := cronet.NewEngineParams()
	engineParams.SetUserAgent("Go-http-client/1.1")

	// Start Cronet engine
	engine := cronet.NewEngine()
	engine.StartWithParams(engineParams)
	engine.StartNetLogToFile("net.log", true)

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
	err := conn.Start(http.MethodGet, os.Args[1], headers, 0, true)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(os.Stdout, conn)
	if err != nil {
		log.Fatal(err)
	}
}
