
# openbgpd-state-server

Serve the structured output of `bgpctl` over HTTP.
The output can then be consumed by the Alice Looking Glass.


## Installation

You will need to have go installed to build the server. Please make
sure your go version is >= 1.10.

Running `go get github.com/alice-lg/openbgp-state-server/cmd/openbgp-state-server`
will give you a binary. You might need to cross-compile
it - GOARCH and GOOS are your friends.

We provide a Makefile for more advanced compilation/configuration.
Running `make static` will create statically linked (linux)
executable.


## Configuration

All runtime configuration is done via commandline flags.

