package main

import (
	"fmt"
	"github.com/miekg/dns"
	"log"
)

const (
	srvName   = "web.service.consul"
	agentAddr = "127.0.0.1:8600"
)

func main() {
	c := new(dns.Client)

	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(srvName), dns.TypeSRV)
	m.RecursionDesired = true

	r, _, err := c.Exchange(m, agentAddr)
	if r == nil {
		log.Fatalf("dns query error: %s\n", err.Error())
	}

	if r.Rcode != dns.RcodeSuccess {
		log.Fatalf("dns query error: %v\n", r.Rcode)
	}

	for _, a := range r.Answer {
		b, ok := a.(*dns.SRV)
		if ok {
			m.SetQuestion(dns.Fqdn(b.Target), dns.TypeA)
			r1, _, err := c.Exchange(m, agentAddr)
			if r1 == nil {
				log.Fatalf("dns query error: %v, %v\n", r1.Rcode, err)
			}
			for _, a1 := range r1.Answer {
				c, ok := a1.(*dns.A)
				if ok {
					fmt.Printf("%s – %s:%d\n", b.Target, c.A, b.Port)
				}
			}
		}
	}
}
