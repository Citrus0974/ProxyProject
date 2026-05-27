package acl

import (
	"encoding/json"
	"os"

	"github.com/Citrus0974/ProxyProject/internal/infrastructure/iptree"
	"github.com/Citrus0974/ProxyProject/internal/models"

	"net/netip"
)

type ACLFile struct {
	IPs   []string `json:"ips"`
	CIDRs []string `json:"cidrs"`
}

type Loader struct {
	allowPath string
	denyPath  string
}

func NewLoader(
	allowPath string,
	denyPath string,
) *Loader {

	return &Loader{
		allowPath: allowPath,
		denyPath:  denyPath,
	}
}

func (l *Loader) Load(
	defaultAllow bool,
) (*models.Rules, error) {

	allowData, err := loadFile(l.allowPath)
	if err != nil {
		return nil, err
	}

	denyData, err := loadFile(l.denyPath)
	if err != nil {
		return nil, err
	}

	rules := &models.Rules{
		AllowIPs: make(map[netip.Addr]struct{}),
		DenyIPs:  make(map[netip.Addr]struct{}),

		AllowTree: iptree.New(),
		DenyTree:  iptree.New(),

		DefaultAllow: defaultAllow,
	}

	for _, ipStr := range allowData.IPs {

		ip, err := netip.ParseAddr(ipStr)
		if err != nil {
			continue
		}

		rules.AllowIPs[ip] = struct{}{}
	}

	for _, ipStr := range denyData.IPs {

		ip, err := netip.ParseAddr(ipStr)
		if err != nil {
			continue
		}

		rules.DenyIPs[ip] = struct{}{}
	}

	for _, cidr := range allowData.CIDRs {
		_ = rules.AllowTree.Add(cidr)
	}

	for _, cidr := range denyData.CIDRs {
		_ = rules.DenyTree.Add(cidr)
	}

	return rules, nil
}

func loadFile(path string) (*ACLFile, error) {

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	result := &ACLFile{}

	err = json.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}