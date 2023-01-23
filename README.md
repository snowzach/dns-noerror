# DNS-NOERROR

What is this? It's a simple DNS server that just returns no error for every request.
It's a hack... you shouldn't probably use it unless you know what you are doing. 
I am using it to make my Mikrotik router happy with local lookups for Alpine
Linux contianers that when receiving NXDOMAIN responses to AAAA records don't
bother to look at A records for the domain.

