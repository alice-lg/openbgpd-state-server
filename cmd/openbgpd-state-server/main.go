package main

import (
	"fmt"

	"github.com/alice-lg/openbgpd-state-server/pkg/server"
)

func main() {
	fmt.Println("openbgpd state server    v.0.1.0")

	addr := "127.0.0.1:29111"

	s := server.StateServer{}
	s.StartHTTP(addr)
}
