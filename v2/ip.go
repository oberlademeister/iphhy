package iphhy

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

// IP is an IP-Address / Mask combo
type IP struct {
	ip   net.IP
	mask int
}

// NewFromNetIP creates an IP from a net.IP
func NewFromNetIP(ip net.IP) *IP {
	ret := &IP{ip: make(net.IP, len(ip))}
	copy(ret.ip, ip)
	ret.MakeHost()
	return ret
}

// Parse creates an IP from a string
func Parse(s string) *IP {
	ip := &IP{}
	if i := strings.LastIndex(s, "/"); i > -1 {
		address := string([]byte(s)[0:i])
		mask := string([]byte(s)[i+1 : len(s)])
		debugln(i, address, mask)
		ip.ip = net.ParseIP(address)
		if ip.ip == nil {
			debugf("ip parse error %s", address)
			return nil
		}
		maskI, err := strconv.Atoi(mask)
		if err != nil {
			return nil
		}
		ip.mask = maskI
		if !MaskOk(ip.ip, maskI) {
			return nil
		}
		return ip
	}
	ip.ip = net.ParseIP(s)
	if ip.ip == nil {
		fmt.Println(ip.ip)
		return nil
	}

	ip.MakeHost()
	return ip
}

// MaskOk returns if the mask is within acceptable range
func MaskOk(ip net.IP, cidr int) bool {
	if cidr < 0 {
		return false
	}
	if v4 := ip.To4(); v4 != nil {
		return cidr <= 32
	}
	return cidr <= 128
}

// MakeHost sets the mask to a host mask
func (i *IP) MakeHost() {
	if v4 := i.ip.To4(); v4 != nil {
		i.mask = 32
	} else {
		i.mask = 128
	}
}
