package models

import (
	"net/netip"

	"github.com/zmap/go-iptree/iptree"
)

type Decision int

const (
	Deny Decision = iota
	Allow
)

type Rules struct {
	AllowIPs map[netip.Addr]struct{}
	DenyIPs  map[netip.Addr]struct{}

	AllowTree *iptree.IPTree
	DenyTree  *iptree.IPTree

	DefaultAllow bool
}
