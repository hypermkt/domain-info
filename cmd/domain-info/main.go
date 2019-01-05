package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/miekg/dns"
)

func main() {
	flag.Parse()
	fqdn := flag.Arg(0)
	fmt.Println("fqdn: " + fqdn)

	// refs: https://github.com/miekg/dns/blob/master/types.go#L2
	qtypes := []uint16{
		dns.TypeNS,
		dns.TypeA,
		dns.TypeMX,
		dns.TypeTXT,
		dns.TypeCNAME,
	}

	config, _ := dns.ClientConfigFromFile("/etc/resolv.conf")
	c := new(dns.Client)
	m := new(dns.Msg)

	for _, qtype := range qtypes {
		m.SetQuestion(dns.Fqdn(flag.Arg(0)), qtype)

		r, _, err := c.Exchange(m, net.JoinHostPort(config.Servers[0], config.Port))
		if r == nil {
			log.Fatalf("*** error: %s\n", err.Error())
		}

		for _, a := range r.Answer {
			fmt.Printf("%v\n", a)
		}
	}
}
