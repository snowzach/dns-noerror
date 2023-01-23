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
			m.Rcode = dns.RcodeSuccess
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
