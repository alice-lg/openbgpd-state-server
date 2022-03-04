
# openbgpd-state-server

Serve the structured output of `bgpctl` over HTTP.
The output can then be consumed by the Alice Looking Glass.


## Installation

You will need to have go installed to build the server. Please make
sure your go version is >= 1.16.

Running `go get github.com/alice-lg/openbgpd-state-server/cmd/openbgpd-state-server`
will give you a binary. You might need to cross-compile
it by passing a GOARCH and GOOS to the make environment.

We provide a Makefile for more advanced compilation/configuration.
Running `make static` will create statically linked (linux)
executable.

You might need to cross-compile it by passing a GOARCH
and GOOS to the make environment.

For example: `GOOS=darwin make static` will produce a static Mac build.
Where `make static` is then just a shorthand for `STATIC=1 GOOS=darwin make`

## Testing

Run the test suite with `make test`.

## Configuration

All runtime configuration is done via commandline flags:

    -l <addr>       Set the listen address  (default: 127.0.0.1:29111)
    -listen <addr>

    -bgpctl "mybgpctl -j -s /path/to/socket"
                    Set the bgpctl command  (default: "bgpctl -j")

    -allow
    -a <pattern>    Allow a command. For example "show neighbor _ timer",
                    
Please note that all commands have to be explicitly allowed.

### Looking Glass Configuration

To use the state server with Alice, you have to allow
the following queries:

```bash
  -a "show neighbor" \
  -a "show summary" \
  -a "show rib neighbor * detail" \
  -a "show rib in detail"
```
