package bgpctl

import (
	"context"
	"os/exec"
)

// BGPCTL is a wrapper for bgpctl with a filter list
// for approved commands.
//
// Additional flags (e.g. -s <socket>) can be added by
// appending to this string.
type BGPCTL struct {
	Name string
	Args []string

	AllowedCommands CommandPatterns
}

// Do runs the configured bgpctl command with
// the request as argumens.
func (ctl *BGPCTL) Do(ctx context.Context, req Request) ([]byte, error) {
	args := append(ctl.Args, req...)
	cmd := exec.CommandContext(ctx, ctl.Name, args...)
	return cmd.CombinedOutput()
}

// DefaultBGPCTL is the bgpctl default command
var DefaultBGPCTL = &BGPCTL{
	Name: "bgpctl",
	Args: []string{"-j"},

	AllowedCommands: DefaultAllowedCommands,
}
