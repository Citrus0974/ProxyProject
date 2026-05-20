package models

import "net/netip"

type Decision int

const (
	Deny Decision = iota
	Allow
)

type Rules struct {
	AllowIPs map[netip.Addr]struct{}
	DenyIPs  map[netip.Addr]struct{}

	AllowCIDRs []netip.Prefix
	DenyCIDRs  []netip.Prefix

	DefaultAllow bool
}

