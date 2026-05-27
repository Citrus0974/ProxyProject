package acl

import (
	"net/netip"
	"sync/atomic"
	"time"

	"github.com/Citrus0974/ProxyProject/internal/infrastructure/ipcache"
	"github.com/Citrus0974/ProxyProject/internal/models"
)

type Service struct {
	rules atomic.Pointer[models.Rules]

	cache ipcache.Cache[string, models.Decision]
}

func NewService(
	rules *models.Rules,
	cacheSize int,
	cacheTTL time.Duration,
) *Service {

	s := &Service{
		cache: ipcache.NewLRU[string, models.Decision](
			cacheSize,
			cacheTTL,
		),
	}

	s.rules.Store(rules)

	return s
}

func (s *Service) ReplaceRules(r *models.Rules) {

	s.rules.Store(r)

	s.cache.Purge()
}

func (s *Service) Check(
	ip netip.Addr,
) models.Decision {

	ipStr := ip.String()

	if decision, ok := s.cache.Get(ipStr); ok {
		return decision
	}

	rules := s.rules.Load()

	if _, ok := rules.DenyIPs[ip]; ok {

		s.cache.Add(ipStr, models.Deny)

		return models.Deny
	}

	if rules.DenyTree.Contains(ip) {

		s.cache.Add(ipStr, models.Deny)

		return models.Deny
	}

	if _, ok := rules.AllowIPs[ip]; ok {

		s.cache.Add(ipStr, models.Allow)

		return models.Allow
	}

	if rules.AllowTree.Contains(ip) {

		s.cache.Add(ipStr, models.Allow)

		return models.Allow
	}

	if rules.DefaultAllow {

		s.cache.Add(ipStr, models.Allow)

		return models.Allow
	}

	s.cache.Add(ipStr, models.Deny)

	return models.Deny
}
