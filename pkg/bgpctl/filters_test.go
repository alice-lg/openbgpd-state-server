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
