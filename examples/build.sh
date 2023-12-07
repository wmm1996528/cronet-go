#!/bin/bash -e

export set CGO_CFLAGS="-I/usr/local/include/cronet"
export set CGO_LDFLAGS="-Wl,-rpath,/usr/local/lib/cronet -L/usr/local/lib/cronet -lcronet"

go build -o ./example1 ./1
go build -o ./example2 ./2
