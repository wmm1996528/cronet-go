package main

import (
	"fmt"
	"github.com/weblifeio/cronet-go"
)

func main() {
	engine := cronet.NewEngine()
	defer engine.Destroy()
	fmt.Println("libcronet " + engine.Version())
	fmt.Println("uad " + engine.DefaultUserAgent())
}
