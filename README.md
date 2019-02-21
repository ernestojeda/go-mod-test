# Go Mod Test

https://travis-ci.org/ernestojeda/go-mod-test.svg?branch=master

Testing what happens with go.sum files between patch versions of Golang

The `go.sum` file in this repository was generated from a Go 1.11.2 base image. If you attempt to use a different base image, you will get checksum mismatches when the go modules are verified against the `go.sum` file.

# Testing

```
$ docker build -t test:1.11.2 -f Dockerfile.golang-1.11.2 .
...truncated output
Successfully tagged test:1.11.2

$ docker build -t test:1.11.5 -f Dockerfile.golang-1.11.5 .
...truncated output

go: verifying github.com/hashicorp/go-rootcerts@v1.0.0: checksum mismatch
	downloaded: h1:Rqb66Oo1X/eSV1x66xbDccZjhJigjg0+e82kpwzSwCI=
	go.sum:     h1:ueI78wUjYExhCvMLow4icJnayNNFRgy0d9EGs/a1T44=
```