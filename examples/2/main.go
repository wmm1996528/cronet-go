package main

import (
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/weblifeio/cronet-go"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type ForwardItem struct {
	Proxy          string            `json:"proxy,omitempty"`
	Method         string            `json:"method,omitempty"`
	Headers        map[string]string `json:"headers,omitempty"`
	Url            string            `json:"url,omitempty"`
	Data           string            `json:"data,omitempty"`
	Timeout        int               `json:"timeout,omitempty"`
	AllowRedirects bool              `json:"allow_redirects,omitempty"`
	Verify         bool              `json:"verify,omitempty"`
	Debug          bool              `json:"debug"`
	RandomTls      bool              `json:"random_tls"`
}

func Forward(c *gin.Context) {
	var forWard ForwardItem
	if err := c.ShouldBindJSON(&forWard); err != nil {
		logrus.Error(err.Error())
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}
	logrus.Infof("开始转发 url: %s  proxy: %s", forWard.Url, forWard.Proxy)

	t1 := time.Now()
	// 开始转发
	client, err := NewClient(forWard, false)
	if err != nil {
		logrus.Error("client初始化失败 ", err.Error())
		c.JSON(400, gin.H{
			"msg": "client初始化失败 " + err.Error(),
		})
		return
	}
	if err := client.Start(); err != nil {
		logrus.Error("请求失败 ", err.Error())
		c.JSON(400, gin.H{
			"msg": "请求失败 " + err.Error(),
		})
		return
	}

	res := gin.H{}
	res["msg"] = "请求成功"
	res["status_code"] = client.GetStatusCode()
	res["proxy"] = forWard.Proxy
	res["headers"] = client.GetResponseHeaders()
	res["cookies"] = client.GetSessionCookies()
	res["cost_time"] = time.Now().Sub(t1).Milliseconds()
	res["url"] = client.GetRespUrl()
	res["text"] = client.GetText()
	defer func() {
		err := client.Response.Body.Close()
		if err != nil {
			fmt.Errorf("text %s", err.Error())
			return
		}
	}()
	client.Client.CloseIdleConnections()
	logrus.Infof("%s 转发 %s %s %d 耗时 %s", client.TlsVersion, forWard.Url, forWard.Proxy, client.GetStatusCode(), time.Now().Sub(t1).String())
	c.JSON(200, res)
}

func NewClient(ward ForwardItem, b bool) {
	engineParams := cronet.NewEngineParams()
	engineParams.SetProxyServer()
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

func main() {
	logrus.SetLevel(logrus.InfoLevel)
	//logrus.SetFormatter(&logrus.TextFormatter{
	//	//HideKeys:        true,
	//	TimestampFormat: "2006-01-02 15:04:05.000", //时间格式
	//	FullTimestamp:   true,haox
	//})
	logrus.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		FieldsOrder:     []string{"type"},
		TimestampFormat: "2006-01-02 15:04:05.000", //时间格式
		NoColors:        true,
	})
	// 写入文件

	//src, err := os.OpenFile(fileName+".log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModeAppend)
	//logrus.SetOutput(src)
	//if err != nil {
	//	fmt.Println("err", err)
	//
	//}

	//logrus.SetOutput(logWriter)
	//logrus.SetOutput(io.MultiWriter(logWriter, os.Stdout))
	logrus.SetOutput(os.Stdout)
	fmt.Println("VERSION 1.0.1")
	r := gin.New() //创建一个默认的路由引擎
	r.POST("/tls/forward", Forward)
	r.Run(":58000")
}
