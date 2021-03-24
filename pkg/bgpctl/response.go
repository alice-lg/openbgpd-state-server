package bgpctl

// This file includes well-known bgpctl responses.

import (
	"time"
)

// NeighborSummary is a summary of the BGP neighbor.
type NeighborSummary struct {
	RemoteAS   string        `json:"remote_as"`
	Group      string        `json:"group"`
	RemoteAddr string        `json:"remote_addr"`
	BGPID      string        `json:"bgpid,omitempty"`
	State      string        `json:"state"`
	LastUpdown time.Duration `json:"last_updown"`
}

// Neighbor is a full neighbor state
type Neighbor struct {
	Summary NeighborSummary
	Config  *NeighborConfig  `json:"config"`
	Stats   *NeighborStats   `json:"stats"`
	Session *NeighborSession `json:"session,omitempty"`
}

// NeighborConfig holds a neighbor configuration
type NeighborConfig struct {
	Template             bool                 `json:"template"`
	Cloned               bool                 `json:"cloned"`
	Passive              bool                 `json:"passive"`
	Down                 bool                 `json:"down"`
	Multihop             bool                 `json:"multihop"`
	MaxPrefix            int                  `json:"max_prefix"`
	MaxPrefixRestart     int                  `json:"max_prefix_restart"`
	TTLSecurity          bool                 `json:"ttl_security"`
	Holdtime             int                  `json:"holdtime"`
	MinHoldtime          int                  `json:"min_holdtime"`
	AnnounceCapabilities bool                 `json:"announce_capabilities"`
	Capabilities         NeighborCapabilities `json:"capabilities"`
}

// NeighborCapabilities TODO
type NeighborCapabilities map[string]interface{}

// NeighborStats are metrics about the neighbor
type NeighborStats struct {
	LastRead  string `json:"last_read"`
	LastWrite string `json:"last_write"`
	Prefixes  struct {
		Sent     int `json:"sent"`
		Received int `json:"received"`
	} `json:"prefixes"`
	Message struct {
		Sent struct {
			Open          int `json:"open"`
			Notifications int `json:"notifications"`
			Updates       int `json:"updates"`
			Keepalives    int `json:"keepalives"`
			RouteRefresh  int `json:"route_refresh"`
			Total         int `json:"total"`
		} `json:"sent"`
		Received struct {
			Open          int `json:"open"`
			Notifications int `json:"notifications"`
			Updates       int `json:"updates"`
			Keepalives    int `json:"keepalives"`
			RouteRefresh  int `json:"route_refresh"`
			Total         int `json:"total"`
		} `json:"received"`
	} `json:"message"`
	Update struct {
		Sent struct {
			Updates   int `json:"updates"`
			Withdraws int `json:"withdraws"`
			Eor       int `json:"eor"`
		} `json:"sent"`
		Received struct {
			Updates   int `json:"updates"`
			Withdraws int `json:"withdraws"`
			Eor       int `json:"eor"`
		} `json:"received"`
	} `json:"update"`
}

// NeighborSession holds session information about the neighbor
type NeighborSession struct {
	Holdtime  int `json:"holdtime"`
	Keepalive int `json:"keepalive"`
	Local     struct {
		Address      string               `json:"address"`
		Port         int                  `json:"port"`
		Capabilities NeighborCapabilities `json:"capabilities"`
	} `json:"local"`
	Remote struct {
		Address      string               `json:"address"`
		Port         int                  `json:"port"`
		Capabilities NeighborCapabilities `json:"capabilities"`
	} `json:"remote"`
	Capabilities NeighborCapabilities `json:"capabilities"`
}
