package bgpctl

// bgpctl is used for interacting with openbgpd.
// This package provides wrappers and command handlers.

// Request wraps an instruction to be run by bgpctl.
// This is being passed to bgpctl as command argument.
type Request []string

// Command returns the command passed to bgpctl
func (req Request) Command() string {
	if len(req) == 0 {
		return ""
	}
	return req[0]
}

// Sanitize removes possibly dangrous input
// from the request. See FilterUnsafeString.
func (req Request) Sanitize() Request {
	req1 := make(Request, 0, len(req))
	for _, arg := range req {
		req1 = append(req1, FilterUnsafeString(arg))
	}
	return req1
}
