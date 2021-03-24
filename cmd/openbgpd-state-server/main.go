package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/alice-lg/openbgpd-state-server/pkg/bgpctl"
	"github.com/alice-lg/openbgpd-state-server/pkg/server"
)

// Flags
var (
	listenAddr  string
	allowedCmds bgpctl.CommandPatterns
	ctl         *bgpctl.BGPCTL
)

// Defaults
var (
	defaultListenAddr = "127.0.0.1:29111"
	defaultCTL        = bgpctl.FromString("bgpctl -j")
)

// Usage
const (
	listenAddrUsage = "server listen address"
	bgpctlUsage     = "bgpctl invocation (default \"bgpctl -j\")"
	allowUsage      = "allow bgpctl command"
)

// Flag Parsers
// -bgpctl
func parseFlagBGPCTL(s string) error {
	ctl = bgpctl.FromString(s)
	return nil
}

// -allow -a
func parseFlagAllow(s string) error {
	allowedCmds.Add(bgpctl.RequestFromString(s))
	return nil
}

func init() {
	// Initialize flags
	ctl = bgpctl.DefaultBGPCTL
	allowedCmds = bgpctl.CommandPatterns{}

	flag.StringVar(&listenAddr, "listen", defaultListenAddr, listenAddrUsage)
	flag.StringVar(&listenAddr, "l", defaultListenAddr, listenAddrUsage+" (shorthand)")
	flag.Func("bgpctl", bgpctlUsage, parseFlagBGPCTL)
	flag.Func("allow", allowUsage, parseFlagAllow)
	flag.Func("a", allowUsage+" (shorthand)", parseFlagAllow)
}

func main() {
	flag.Parse()

	// Runtime information
	fmt.Printf("openbgpd state server @ %s\t\tv.%s (%s)\n",
		listenAddr,
		server.Version,
		server.Build)

	if len(allowedCmds) == 0 {
		fmt.Fprintf(os.Stderr, "error: no allowed commands\n")
		fmt.Fprintf(os.Stderr, "please explicitly allow commands using the -allow flag\n")
		os.Exit(-1)
	}

	fmt.Println("allowed commands:")
	for _, cmd := range allowedCmds {
		fmt.Println("  -", cmd)
	}
	ctl.AllowedCommands = allowedCmds

	// Start the server
	s := server.StateServer{BGPCTL: ctl}
	s.StartHTTP(listenAddr)
}
