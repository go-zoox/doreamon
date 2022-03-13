package dns

import (
	"time"

	"github.com/miekg/dns"
)

type Client struct {
	server string
}

const DEFAULT_DNS_SERVER = "114.114.114.114"

func (client Client) LookUp(domain string) ([]string, error) {
	dnsServer := client.server
	if dnsServer == "" {
		dnsServer = DEFAULT_DNS_SERVER
	}

	c := dns.Client{
		Timeout: 5 * time.Second,
	}

	m := dns.Msg{}
	m.SetQuestion(domain+".", dns.TypeA)
	r, _, err := c.Exchange(&m, dnsServer+":53")
	if err != nil {
		return nil, err
	}

	dst := []string{}
	for _, answer := range r.Answer {
		record, ok := answer.(*dns.A)
		if ok {
			dst = append(dst, record.A.String())
		}
	}

	return dst, nil
}
