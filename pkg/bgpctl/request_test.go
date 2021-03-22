package bgpctl

import (
	"testing"
)

func TestRequestFromString(t *testing.T) {
	req := RequestFromString("show neighbor 2001:504:2f::852:1")
	if req[0] != "show" && req[1] != "neighbor" && req[2] != "2001:504:2f::852:1" {
		t.Error("unexpected result", req)
	}
}
