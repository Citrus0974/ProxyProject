package models

import (
	"net/netip"

	"github.com/Citrus0974/ProxyProject/internal/infrastructure/iptree"
)

type Decision int

const (
	Deny Decision = iota
	Allow
)

type Rules struct {
	AllowIPs map[netip.Addr]struct{}
	DenyIPs  map[netip.Addr]struct{}

	AllowTree iptree.Tree
	DenyTree  iptree.Tree

	DefaultAllow bool
}
