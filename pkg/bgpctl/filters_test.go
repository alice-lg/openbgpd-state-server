package bgpctl

import (
	"testing"
)

func TestFilterUnsafeString(t *testing.T) {
	allowed := "abcdefghijklmnopqrstuvwxyzABCDETC ..:0:.. /"
	filtered := FilterUnsafeString(allowed)
	if allowed != filtered {
		t.Error("expected filtered to be identical to allowed")
	}

	badUTF8 := "`${}()test--𝕯٤ḞԍНǏ𝙅ƘԸⲘ𝙉`০Ρ𝗤Ɍ𝓢ȚЦ𝒱Ѡ𝓧Ƴ"
	filtered = FilterUnsafeString(badUTF8)
	if filtered != "test" {
		t.Error("filtered should only allow 'test'")
	}
}

func TestAllowPatterns(t *testing.T) {
	p := AllowPatterns{
		Request{"foo", "*"},
		Request{"bar"}}

	if !p.IsAllowed(Request{"foo", "23.42.123.5"}) {
		t.Error("expected this request to be allowed")
	}

	if !p.IsAllowed(Request{"bar"}) {
		t.Error("this request should be allowed")
	}

	if p.IsAllowed(Request{"bar", "baz"}) {
		t.Error("this request should not be allowed")
	}

	if p.IsAllowed(Request{"foo"}) {
		t.Error("this request should not be allowed")
	}
}
