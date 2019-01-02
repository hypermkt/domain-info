package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {
	flag.Parse()
	domain := flag.Arg(0)
	fmt.Println("fqdn: " + domain)

	// NS
	nss, error := net.LookupNS(domain)
	if error != nil {
		// log error
	}
	for _, ns := range nss {
		fmt.Printf("Name Server: %v\n", ns.Host)
	}

	// A
	ips, error := net.LookupIP(domain)
	if error != nil {
		// log error
	}
	for _, ip := range ips {
		fmt.Printf("A: %v\n", ip)
	}

	// MX
	mxs, error := net.LookupMX(domain)
	if error != nil {
		// log error
	}
	for _, mx := range mxs {
		fmt.Printf("MX: %v\n", mx.Host)
	}

	// TXT
	txts, error := net.LookupTXT(domain)
	if error != nil {
		// log error
	}
	for _, txt := range txts {
		fmt.Printf("TXT: %v\n", txt)
	}
}
