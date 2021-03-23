package bgpctl

import (
	"regexp"
)

// RegMatchUnsafe is a regular expression matching
// not a-z, A-Z, 0-9 or :, / and ' '
var RegMatchUnsafe = regexp.MustCompile(`[^a-zA-Z0-9:\.\s\/]+`)

// FilterUnsafeString removes anything not alphanumeric
// from the string.
func FilterUnsafeString(s string) string {
	return RegMatchUnsafe.ReplaceAllString(s, "")
}

// CommandPatterns is a list of commands that
// are allowed to be run. The '*' wildcard character
// marks for any allowed string. For example
//   show neighbor * timers
// would allow for querying timers
type CommandPatterns [][]string

// IsAllowed checks if the command matches the pattern
func (p CommandPatterns) IsAllowed(req Request) bool {
	for _, pattern := range p {
		if len(pattern) != len(req) {
			continue // this can not match
		}
		found := true
		for i, t := range req {
			if pattern[i] == "*" {
				continue // still match
			}
			if pattern[i] != t {
				found = false
				break // not a match
			}
		}
		if found {
			return true
		}
	}
	return false
}
