package test

import (
	"testing"
	"time"

	"github.com/miekg/dns"
)

func TestProxyThreeWay(t *testing.T) {
	// Run 3 DNServer server, 2 upstream ones and a proxy. 1 Upstream is unhealthy after 1 query, but after
	// that we should still be able to send to the other one.

	// Backend DNServer's.
	corefileUp1 := `example.org:0 {
		erratic {
			drop 2
		}
	}`

	up1, err := DNServerServer(corefileUp1)
	if err != nil {
		t.Fatalf("Could not get DNServer serving instance: %s", err)
	}
	defer up1.Stop()

	corefileUp2 := `example.org:0 {
		whoami
	}`

	up2, err := DNServerServer(corefileUp2)
	if err != nil {
		t.Fatalf("Could not get DNServer serving instance: %s", err)
	}
	defer up2.Stop()

	addr1, _ := DNServerServerPorts(up1, 0)
	if addr1 == "" {
		t.Fatalf("Could not get UDP listening port")
	}
	addr2, _ := DNServerServerPorts(up2, 0)
	if addr2 == "" {
		t.Fatalf("Could not get UDP listening port")
	}

	// Proxying DNServer.
	corefileProxy := `example.org:0 {
		forward . ` + addr1 + " " + addr2 + ` {
			max_fails 1
		}
	}`

	prx, err := DNServerServer(corefileProxy)
	if err != nil {
		t.Fatalf("Could not get DNServer serving instance: %s", err)
	}
	defer prx.Stop()
	addr, _ := DNServerServerPorts(prx, 0)
	if addr == "" {
		t.Fatalf("Could not get UDP listening port")
	}

	m := new(dns.Msg)
	m.SetQuestion("example.org.", dns.TypeA)
	c := new(dns.Client)
	c.Timeout = 10 * time.Millisecond

	for i := 0; i < 10; i++ {
		r, _, err := c.Exchange(m, addr)
		if err != nil {
			continue
		}
		// We would previously get SERVFAIL, so just getting answers here
		// is a good sign. The actual timeouts are handled in the err != nil case
		// above.
		if r.Rcode != dns.RcodeSuccess {
			t.Fatalf("Expected success rcode, got %d", r.Rcode)
		}
	}
}
