package iptree

import (
	"net/netip"

	goiptree "github.com/zmap/go-iptree/iptree"
)

type Tree interface {
	Add(cidr string) error
	Contains(ip netip.Addr) bool
}

type IPTree struct {
	tree *goiptree.IPTree
}

func New() *IPTree {
	return &IPTree{
		tree: goiptree.New(),
	}
}

func (t *IPTree) Add(cidr string) error {
	return t.tree.AddByString(cidr, true)
}

func (t *IPTree) Contains(ip netip.Addr) bool {
	_, ok, _ := t.tree.GetByString(ip.String())
	return ok
}