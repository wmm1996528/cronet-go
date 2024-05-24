#!/bin/bash -e

#export set CGO_CFLAGS="-I"
export set CGO_CFLAGS="-I/Users/wang/Desktop/cronet-binaries/src/out/Release/cronet"
#export set CGO_LDFLAGS="-Wl,-rpath, -L -lcronet"
export set CGO_LDFLAGS="-Wl,-rpath,/Users/wang/Desktop/cronet-binaries/src/out/Release/cronet -L/Users/wang/Desktop/cronet-binaries/src/out/Release/cronet -lcronet"
export GOPROXY=https://goproxy.cn
go build -o ./example1 ./1
#go build -o ./example2 ./2
#go build -o ./example3 ./3
