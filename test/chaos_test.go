package test

import (
	"testing"

	// Plug in DNServer, needed for AppVersion and AppName in this test.
	"github.com/coredns/caddy"
	_ "github.com/khulnasoft-lab/dnserver/coremain"

	"github.com/miekg/dns"
)

func TestChaos(t *testing.T) {
	corefile := `.:0 {
		chaos
	}`

	i, udp, _, err := DNServerServerAndPorts(corefile)
	if err != nil {
		t.Fatalf("Could not get DNServer serving instance: %s", err)
	}
	defer i.Stop()

	m := new(dns.Msg)
	m.SetQuestion("version.bind.", dns.TypeTXT)
	m.Question[0].Qclass = dns.ClassCHAOS

	resp, err := dns.Exchange(m, udp)
	if err != nil {
		t.Fatalf("Expected to receive reply, but didn't: %v", err)
	}
	chTxt := resp.Answer[0].(*dns.TXT).Txt[0]
	version := caddy.AppName + "-" + caddy.AppVersion
	if chTxt != version {
		t.Fatalf("Expected version to be %s, got %s", version, chTxt)
	}
}
