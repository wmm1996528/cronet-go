#!/bin/bash -e

export set CGO_CFLAGS="-I/Users/wang/Desktop/cronet-binaries/src/out/Release/cronet"
export set CGO_LDFLAGS="-Wl,-rpath,/Users/wang/Desktop/cronet-binaries/src/out/Release/cronet -L/Users/wang/Desktop/cronet-binaries/src/out/Release/cronet -lcronet"

go build -o ./example1 ./1
go build -o ./example2 ./2
