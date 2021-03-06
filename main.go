package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
)

func main() {
	var target ip
	flag.Var(&target, "target", "target IP address")
	flag.Parse()

	resp, err := http.Get("https://ip-ranges.amazonaws.com/ip-ranges.json")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Fatal(fmt.Errorf("(%d) %s", resp.StatusCode, body))
	}

	var r ranges
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		log.Fatal(err)
	}

	for _, prefix := range append(r.Prefixes, r.IPv6Prefixes...) {
		addr := prefix.IPPrefix
		if addr == "" {
			addr = prefix.IPv6Prefix
		}
		_, cidr, err := net.ParseCIDR(addr)
		if err != nil {
			log.Printf("error parsing CIDR (%s): %s", addr, err)
			continue
		}

		if cidr.Contains(net.IP(target)) {
			fmt.Println("Target: ", target)
			fmt.Println(prefix)
			break
		}
	}
}

type ranges struct {
	SyncToken    string   `json:"syncToken"`
	CreateDate   string   `json:"createDate"`
	Prefixes     []prefix `json:"prefixes"`
	IPv6Prefixes []prefix `json:"ipv6_prefixes"`
}

type prefix struct {
	IPPrefix           string `json:"ip_prefix,omitempty"`
	IPv6Prefix         string `json:"ipv6_prefix,omitempty"`
	Region             string `json:"region"`
	NetworkBorderGroup string `json:"network_border_group"`
	Service            string `json:"service"`
}

func (p prefix) String() string {
	var b strings.Builder

	b.WriteString("Prefix:  ")
	if p.IPPrefix != "" {
		b.WriteString(p.IPPrefix)
	}

	if p.IPv6Prefix != "" {
		b.WriteString(p.IPv6Prefix)
	}

	b.WriteString("\nRegion:  ")
	b.WriteString(p.Region)

	b.WriteString("\nNBG:     ")
	b.WriteString(p.NetworkBorderGroup)

	b.WriteString("\nService: ")
	b.WriteString(p.Service)

	return b.String()
}

type ip net.IP

func (v *ip) Set(value string) error {
	*v = ip(net.ParseIP(value))

	return nil
}

func (v ip) String() string {
	return net.IP(v).String()
}
