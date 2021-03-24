package bgpctl

import (
	"testing"
)

func TestFromString(t *testing.T) {
	ctl := FromString("bgpctl")
	if ctl.Name != "bgpctl" {
		t.Error("unexpected name:", ctl.Name)
	}
	if len(ctl.Args) != 0 {
		t.Error("ctl should not have args")
	}

	ctl = FromString("bgpctl -j")
	if len(ctl.Args) != 1 {
		t.Error("ctl should have 1 arg")
	}
}
