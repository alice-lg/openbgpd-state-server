package bgpctl

import (
	"context"
	"errors"
	"os/exec"
)

// ErrCommandDisallowed will be returned if the
// request does not match any pattern in AllowedCommands
var ErrCommandDisallowed = errors.New("the requested command is not approved")

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

// FromString parses the bgpctl invocation
// and returns a new wrapped BGPCTL.
func FromString(s string) *BGPCTL {
	req := RequestFromString(s)
	ctl := &BGPCTL{
		Name:            req.Command(),
		Args:            req.Args(),
		AllowedCommands: CommandPatterns{},
	}
	return ctl
}

// Do runs the configured bgpctl command with
// the request as argumens.
func (ctl *BGPCTL) Do(ctx context.Context, req Request) ([]byte, error) {
	if !ctl.AllowedCommands.IsAllowed(req) {
		return nil, ErrCommandDisallowed
	}

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
