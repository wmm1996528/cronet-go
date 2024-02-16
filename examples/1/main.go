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
	engineParams.SetUserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	engineParams.SetProxyServer(proxyArg.String())
	//engineParams.SetEnableHTTP2(true)
	//engineParams.SetEnableQuic(false)
	t := cronet.NewCronetTransport(engineParams, true)
	//ConfigureClientCertificate(&t.Engine, certPath, keyPath, []string{urlArg.Host, proxyArg.Host})

	client := &http.Client{
		Transport: t,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, _ := http.NewRequest("GET", "https://booking.jetstar.com/au/en/booking/search-flights?s=true&adults=1&children=0&infants=0&selectedclass1=economy&currency=AUD&mon=true&channel=DESKTOP&origin1=ADL&destination1=BNE&departuredate1=2024-03-13", nil)

	req.Header.Set("Host", "booking.jetstar.com")
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Set("authority", "www.jetstar.com")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("content-type", "text/plain;charset=UTF-8")
	req.Header.Set("origin", "https://www.jetstar.com")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 14_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.39(0x18002733) NetType/WIFI Language/zh_CN")
	req.Header.Set("cookie", "bm_sz=898E219CCD0C9268C8DFB3DF49B3CA73~YAAQ9KzbF5iNfhKOAQAAlgQKHReb++txKDS8G4xMSUgZSOhkfqqdqdaYynorNP94AFC/7ksC/+W03WpzFWDBZxLEFMCMDFNnwXUb+q9kx/+0pZEsunoYNe+AWuD0Ub/3oljposvOdn0ZWvcvcvLv7uvwb8ZNNsD74txBv3UdzcszoZk0HTN/KvCPQc1JoYZX0RP4grOv6dkWrOSX0cHogv5pnC8FLkeg/nUx7wy1xALDXhpcaBe4hytV6jiTnqhxC4syOVVgquPcvkcsQIYr3iXkv6hjZ6RbmRWwW9sOwGUc4bhXIre82zs8P1YdhAa/V3Xet3UGm5ksmL/+iXaFeITAm6ZI2dqq+IfprjBiGwjNaLabqsRHAQ==~3425329~4473670; ak_bmsc=BA6564EB64C4C49733F0A37904B9DDEF~000000000000000000000000000000~YAAQ9KzbF5aNfhKOAQAAlgQKHReFVZ24wM6VCDE4CB1isfAyg4vOiPTNgpaF0FTiiHU2ZFFepcdPp6vdGAWRrEbOn+ZDG+BVNJymeSmLPb/AMHxYnhTbKELlFxoYwmXqoOZf13oF40TLPdtTyk16sL8TiOgCqVfQZGuX80e6m/JosxKJxiIyr+wgiepYVWei3PSzTSABXkXu9MxO9VCM0Rs/t6Dxof8udtXWCvPNrrfEIjuEJfB/KKdH4gOg9ZSyIGHd6W9r/jmP+CWoW6Cx32Vme3ZGVUWdz1F+TdtWQmax2FkcnJM15Y0/XEqImUWz53IetDP0oQsUTRPA66JkYJCJoL2SPtJ4jEq4zSje+cHfD5gB0aVjkqd97k8H22ODHsKwUBDoIN69onTU; _abck=6A851BB2D9F247BF60F3845E6EB77B03~-1~YAAQ9KzbF6GNfhKOAQAAdg8KHQvfV361NYnaRNxAx89ZVsvG5Noq0PLODY4MA4RUezktH+1CkhjOHQ5vB7k0Xo7tnhjOSrVXUnUyEx7PtM7TJoTKZ7mEPRz/Oec+pmJ1X1mYGCXTYQH2mQv0TD9ixKvrcrcNB3IRdMGAvMkQ/6y6HSX0sxl6LC0FRhhJuaG6iCSsDFcqfB0V8+mQyjXjgHg5Z6YOVlX1HWcZR3YtL886X8k0tvA/eiPu93O/nFdyNGY6L+mQSIH/Mde2tGLO0mCFMX+uFwEjATNZmhpl5qiUv+kKYjo9jQysPKqe1Ty4Xm18cT8lcMfJoq38r+V1dzMoE+pmRGWe5uBJtCu7DHBooNAWuVCmq4jyBQPfqNUP/EXzoP7YQoQ+ExgN~-1~-1~-1; user-location=\"country_code=US,region_code=CO,city=BOULDER,lat=40.0440,long=-105.194\"")

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	//resp, err := client.Get(urlArg.String())
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}
