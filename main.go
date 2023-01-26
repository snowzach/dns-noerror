package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/miekg/dns"
)

var port = flag.Int("p", 5353, "port number")

func main() {
	flag.Parse()

	// Handle every query request with a success and no response
	dns.HandleFunc(".", func(w dns.ResponseWriter, r *dns.Msg) {
		m := new(dns.Msg)
		m.SetReply(r)
		m.Compress = false
		switch r.Opcode {
		case dns.OpcodeQuery:
			if len(r.Question) > 0 && r.Question[0].Qtype == dns.TypeAAAA {
				// Effectively return no answer for AAAA lookups
				m.Rcode = dns.RcodeSuccess
			} else {
				m.Rcode = dns.RcodeNameError
			}

		}
		w.WriteMsg(m)
	})

	// start server
	server := &dns.Server{Addr: ":" + strconv.Itoa(*port), Net: "udp"}
	log.Printf("Listening on :%d\n", *port)
	err := server.ListenAndServe()
	defer server.Shutdown()
	if err != nil {
		log.Fatalf("Failed to start server: %v\n ", err)
	}
}
