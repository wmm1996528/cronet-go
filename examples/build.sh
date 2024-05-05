#!/bin/bash -e

export set CGO_CFLAGS="-I/home/wang/cronet-go/examples/cron"
export set CGO_LDFLAGS="-Wl,-rpath,/home/wang/cronet-go/examples/cron -L/home/wang/cronet-go/examples/cron -lcronet"
export GOPROXY=https://goproxy.cn
go build -o ./example1 ./1
go build -o ./example2 ./2
go build -o ./example3 ./3
