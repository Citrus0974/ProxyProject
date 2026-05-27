package acl_test

import (
	"net/netip"
	"testing"
	"time"

	"github.com/Citrus0974/ProxyProject/internal/infrastructure/iptree"
	"github.com/Citrus0974/ProxyProject/internal/models"
	"github.com/Citrus0974/ProxyProject/internal/usecase/acl"

	"github.com/stretchr/testify/require"
)

func buildRules() *models.Rules {

	allowTree := iptree.New()
	denyTree := iptree.New()

	_ = allowTree.Add("10.0.0.0/8")
	_ = denyTree.Add("192.168.0.0/16")

	return &models.Rules{
		AllowIPs: map[netip.Addr]struct{}{
			netip.MustParseAddr("127.0.0.1"): {},
		},

		DenyIPs: map[netip.Addr]struct{}{
			netip.MustParseAddr("8.8.8.8"): {},
		},

		AllowTree: allowTree,
		DenyTree:  denyTree,

		DefaultAllow: false,
	}
}

func TestAllowIP(t *testing.T) {

	svc := acl.NewService(
		buildRules(),
		100,
		time.Minute,
	)

	decision := svc.Check(
		netip.MustParseAddr("127.0.0.1"),
	)

	require.Equal(
		t,
		models.Allow,
		decision,
	)
}

func TestDenyIP(t *testing.T) {

	svc := acl.NewService(
		buildRules(),
		100,
		time.Minute,
	)

	decision := svc.Check(
		netip.MustParseAddr("8.8.8.8"),
	)

	require.Equal(
		t,
		models.Deny,
		decision,
	)
}

func TestAllowCIDR(t *testing.T) {

	svc := acl.NewService(
		buildRules(),
		100,
		time.Minute,
	)

	decision := svc.Check(
		netip.MustParseAddr("10.1.1.1"),
	)

	require.Equal(
		t,
		models.Allow,
		decision,
	)
}

func TestDenyCIDR(t *testing.T) {

	svc := acl.NewService(
		buildRules(),
		100,
		time.Minute,
	)

	decision := svc.Check(
		netip.MustParseAddr("192.168.1.1"),
	)

	require.Equal(
		t,
		models.Deny,
		decision,
	)
}

func TestDefaultPolicy(t *testing.T) {

	svc := acl.NewService(
		buildRules(),
		100,
		time.Minute,
	)

	decision := svc.Check(
		netip.MustParseAddr("1.1.1.1"),
	)

	require.Equal(
		t,
		models.Deny,
		decision,
	)
}