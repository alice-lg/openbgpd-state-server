package server

var (
	// Version is the server version. Usually the git tag
	// pointing to HEAD. See top level Makefile for details.
	Version string = "HEAD"

	// Build is the server build id. Usually the
	// short git revision hash of HEAD.
	Build string = "unknown"
)

// Status describes the server status. It contains
// a bgpd version (if identifiable), a state server version
// and the uptime. See server.Version and Build for details.
type Status struct {
	Service string `json:"service"`
	Version string `json:"version"`
	Build   string `json:"build"`
}
