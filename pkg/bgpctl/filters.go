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
