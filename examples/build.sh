#!/bin/bash -e

export set CGO_CFLAGS="-I/Users/wang/Desktop/cronet-go/examples/cron"
export set CGO_LDFLAGS="-Wl,-rpath,/Users/wang/Desktop/cronet-go/examples/cron -L/Users/wang/Desktop/cronet-go/examples/cron -lcronet"

go build -o ./example1 ./1
go build -o ./example2 ./2
