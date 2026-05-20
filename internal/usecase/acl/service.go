package acl

import (
	"net/netip"
	"sync/atomic"

	"github.com/Citrus0974/ProxyProject/internal/models"
)

type Service struct {
	rules atomic.Pointer[models.Rules]
}

func NewService(rules *models.Rules) *Service {
	s := &Service{}
	s.rules.Store(rules)
	return s
}

func (s *Service) ReplaceRules(r *models.Rules) {
	s.rules.Store(r)
}

func (s *Service) Check(ip netip.Addr) models.Decision {

	rules := s.rules.Load()

	if _, ok := rules.DenyIPs[ip]; ok {
		return models.Deny
	}

	for _, prefix := range rules.DenyCIDRs {
		if prefix.Contains(ip) {
			return models.Deny
		}
	}

	if _, ok := rules.AllowIPs[ip]; ok {
		return models.Allow
	}

	for _, prefix := range rules.AllowCIDRs {
		if prefix.Contains(ip) {
			return models.Allow
		}
	}

	if rules.DefaultAllow {
		return models.Allow
	}

	return models.Deny
}
