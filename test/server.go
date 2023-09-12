package test

import (
	"sync"

	"github.com/coredns/caddy"
	_ "github.com/khulnasoft-lab/dnserver/core" // Hook in DNServer.
	"github.com/khulnasoft-lab/dnserver/core/dnsserver"
	_ "github.com/khulnasoft-lab/dnserver/core/plugin" // Load all managed plugins in github.com/khulnasoft-lab/dnserver.
)

var mu sync.Mutex

// DNServerServer returns a DNServer test server. It just takes a normal Corefile as input.
func DNServerServer(corefile string) (*caddy.Instance, error) {
	mu.Lock()
	defer mu.Unlock()
	caddy.Quiet = true
	dnsserver.Quiet = true

	return caddy.Start(NewInput(corefile))
}

// DNServerServerStop stops a server.
func DNServerServerStop(i *caddy.Instance) { i.Stop() }

// DNServerServerPorts returns the ports the instance is listening on. The integer k indicates
// which ServerListener you want.
func DNServerServerPorts(i *caddy.Instance, k int) (udp, tcp string) {
	srvs := i.Servers()
	if len(srvs) < k+1 {
		return "", ""
	}
	u := srvs[k].LocalAddr()
	t := srvs[k].Addr()

	if u != nil {
		udp = u.String()
	}
	if t != nil {
		tcp = t.String()
	}
	return
}

// DNServerServerAndPorts combines DNServerServer and DNServerServerPorts to start a DNServer
// server and returns the udp and tcp ports of the first instance.
func DNServerServerAndPorts(corefile string) (i *caddy.Instance, udp, tcp string, err error) {
	i, err = DNServerServer(corefile)
	if err != nil {
		return nil, "", "", err
	}
	udp, tcp = DNServerServerPorts(i, 0)
	return i, udp, tcp, nil
}

// Input implements the caddy.Input interface and acts as an easy way to use a string as a Corefile.
type Input struct {
	corefile []byte
}

// NewInput returns a pointer to Input, containing the corefile string as input.
func NewInput(corefile string) *Input {
	return &Input{corefile: []byte(corefile)}
}

// Body implements the Input interface.
func (i *Input) Body() []byte { return i.corefile }

// Path implements the Input interface.
func (i *Input) Path() string { return "Corefile" }

// ServerType implements the Input interface.
func (i *Input) ServerType() string { return "dns" }
