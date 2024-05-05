package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/sagernet/sing/common/bufio"
	E "github.com/sagernet/sing/common/exceptions"
	M "github.com/sagernet/sing/common/metadata"
	N "github.com/sagernet/sing/common/network"
	"github.com/sagernet/sing/common/redir"
	"github.com/sagernet/sing/transport/mixed"
	"github.com/weblifeio/cronet-go"
	"log"
	"net"
	"net/netip"
	"net/url"
	"os"
	"os/signal"
	"syscall"
)

const kFirstPaddings = 8

type Listener struct {
	url           string
	authorization string
	engine        cronet.StreamEngine
	extraHeaders  map[string]string
}

func (l *Listener) NewConnection(ctx context.Context, conn net.Conn, metadata M.Metadata) error {
	log.Println(metadata.Source, " => ", metadata.Destination)
	headers := map[string]string{
		"-connect-authority": metadata.Destination.String(),
		//"Padding":            generatePaddingHeader(),
	}
	if l.authorization != "" {
		headers["proxy-authorization"] = l.authorization
	}
	for key, value := range l.extraHeaders {
		headers[key] = value
	}
	bidiConn := l.engine.CreateConn(true, false)
	err := bidiConn.Start("CONNECT", l.url, headers, 0, false)
	fmt.Println(l.url, headers)
	if err != nil {
		return E.Cause(err, "start bidi conn")
	}
	return bufio.CopyConn(ctx, conn, bidiConn)
}

func (l *Listener) NewPacketConnection(ctx context.Context, conn N.PacketConn, metadata M.Metadata) error {
	conn.Close()
	return nil
}

func (l *Listener) HandleError(err error) {
	if E.IsClosed(err) {
		return
	}
	log.Fatal(err)
}
func main() {
	err := os.Setenv("GODEBUG", "invalidptr=0")
	if err != nil {
		panic(err)
	}
	engine := cronet.NewEngine()
	params := cronet.NewEngineParams()
	//proxyURL, err := url.Parse("http://user-uni003-region-de-sessid-1125-sesstime-5-keep-true:q39CEBTs5A5YQXor@pr.roxlabs.cn:4600")
	proxyURL, err := url.Parse("https://admin:123456@localhost:8081")
	fmt.Println(proxyURL.String())
	if err != nil {
		panic(err)
	}
	switch proxyURL.Scheme {
	case "https":
	case "http":
		params.SetEnableHTTP2(false)
		params.SetEnableQuic(false)
	case "quic":
		params.SetEnableHTTP2(false)
		params.SetEnableQuic(true)
	default:
	}
	params.SetEnableHTTP2(false)
	params.SetEnableQuic(false)
	fmt.Println(params)
	engine.StartNetLogToFile("./1.log", true)

	var proxyAuthorization string
	if proxyURL.User != nil {
		password, _ := proxyURL.User.Password()
		proxyAuthorization = "Basic " + base64.StdEncoding.EncodeToString([]byte(proxyURL.User.Username()+":"+password))
		proxyURL.User = nil
	}

	engine.StartWithParams(params)
	params.Destroy()
	listener := &Listener{
		url:           proxyURL.String(),
		authorization: proxyAuthorization,
		engine:        engine.StreamEngine(),
		extraHeaders:  make(map[string]string),
	}
	bind, err := netip.ParseAddrPort("127.0.0.1:9991")
	var transMode redir.TransproxyMode

	inbound := mixed.NewListener(bind, nil, transMode, 300, listener)
	err = inbound.Start()
	if err != nil {
		log.Fatal(err)
	}

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)
	<-osSignals

	engine.Shutdown()
	engine.Destroy()
	inbound.Close()

}
