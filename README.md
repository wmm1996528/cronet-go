# cronet-go

Cronet is the Chromium network stack made available as a library. Cronet takes advantage of multiple
technologies that reduce the latency and increase the throughput of the network requests.

The Cronet Library handles the requests of apps used by millions of people on a daily basis, such as YouTube, Google
App, Google Photos, and Maps - Navigation & Transit.

This experimental project ported Cronet to golang. To learn how to use the Cronet Library in
your app, see the [examples](./examples).

## Build Cronet Library

Follow all the [Get the Code](https://www.chromium.org/developers/how-tos/get-the-code/) instructions for your target platform up to and including running hooks.

Apply weblifeio customization:

```sh
git remote add weblifeio https://github.com/weblifeio/chromium
git fetch weblifeio
git cherry-pick weblifeio/develop ^weblifeio/main
```

Follow the [instructions](https://chromium.googlesource.com/chromium/src/+/master/components/cronet/build_instructions.md#desktop-builds-targets-the-current-os) for Desktop builds.

## Install Cronet Library

The following instructions assume you have switched to the `chromium/src` directory:

```sh
mkdir /usr/local/include/cronet
mkdir /usr/local/lib/cronet

cp out/Cronet/cronet/include/* /usr/local/include/cronet
cp out/Cronet/*.dylib /usr/local/lib/cronet

CRONET_VERSION=$(build/util/version.py -f out/Cronet/cronet/VERSION -t "@MAJOR@.@MINOR@.@BUILD@.@PATCH@") \
&& ln -sf /usr/local/lib/cronet/libcronet.${CRONET_VERSION}.dylib /usr/local/lib/cronet/libcronet.124.0.6344.0.dylib \
&& ln -sf /usr/local/lib/cronet/libcronet.${CRONET_VERSION}.dylib /usr/local/lib/libcronet.124.0.6344.0.dylib \
&& ln -sf /usr/local/lib/cronet/libcronet.${CRONET_VERSION}.dylib /usr/local/lib/libcronet.${CRONET_VERSION}.dylib
```

Replace `.dylib` to `.so` if you're on Linux.

## Use cronet-go

```
import (
	"github.com/weblifeio/cronet-go"
)
```

When building your project set these environment variables:

```sh
export set CGO_CFLAGS="-I/usr/local/include/cronet"
export set CGO_LDFLAGS="-Wl,-rpath,/usr/local/lib/cronet -L/usr/local/lib/cronet -lcronet"

go build <your-project-here>
```
