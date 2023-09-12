package minimal

import (
	"github.com/coredns/caddy"
	"github.com/khulnasoft-lab/dnserver/core/dnsserver"
	"github.com/khulnasoft-lab/dnserver/plugin"
)

func init() {
	plugin.Register("minimal", setup)
}

func setup(c *caddy.Controller) error {
	c.Next()
	if c.NextArg() {
		return plugin.Error("minimal", c.ArgErr())
	}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		return &minimalHandler{Next: next}
	})

	return nil
}
