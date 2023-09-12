package local

import (
	"github.com/coredns/caddy"
	"github.com/khulnasoft-lab/dnserver/core/dnsserver"
	"github.com/khulnasoft-lab/dnserver/plugin"
)

func init() { plugin.Register("local", setup) }

func setup(c *caddy.Controller) error {
	l := Local{}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		l.Next = next
		return l
	})

	return nil
}
