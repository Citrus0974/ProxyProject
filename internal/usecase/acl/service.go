package acl

import (
	"net/netip"
	"sync/atomic"
	"time"

	"github.com/Citrus0974/ProxyProject/internal/models"
	lru "github.com/hashicorp/golang-lru/v2/expirable"
)

type Service struct {
	rules atomic.Pointer[models.Rules]
	cache *lru.LRU[netip.Addr, models.Decision]
}

func NewService(rules *models.Rules, cacheSize int, cacheTTL time.Duration) *Service {
	s := &Service{}
	s.rules.Store(rules)
	s.cache = lru.NewLRU[netip.Addr, models.Decision](cacheSize, nil, cacheTTL)
	return s
}

func (s *Service) ReplaceRules(r *models.Rules) {
	s.rules.Store(r)
	s.cache.Purge()
}

func (s *Service) Check(ip netip.Addr) models.Decision {
	ipStr := ip.String()

	if decision, ok := s.cache.Get(ip); ok {
		return decision
	}

	rules := s.rules.Load()

	if _, ok := rules.DenyIPs[ip]; ok {
		s.cache.Add(ip, models.Deny)
		return models.Deny
	}

	if _, ok, _ := rules.DenyTree.GetByString(ipStr); ok {
		s.cache.Add(ip, models.Deny)
		return models.Deny
	}

	if _, ok := rules.AllowIPs[ip]; ok {
		return models.Allow
	}

	if _, ok, _ := rules.AllowTree.GetByString(ipStr); ok {
		s.cache.Add(ip, models.Allow)
		return models.Allow
	}

	if rules.DefaultAllow {
		s.cache.Add(ip, models.Allow)
		return models.Allow
	}

	return models.Deny
}
