package iphhy

import (
	"fmt"
	"math/big"
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
	ret.MakeHostInplace()
	return ret
}

// NewFromBigInt creates an IP from a net.IP
func NewFromBigInt(i *big.Int) *IP {
	b := i.Bytes()
	ret := &IP{
		ip: make(net.IP, net.IPv6len),
	}
	if len(b) > net.IPv6len {
		return nil
	}
	if len(b) <= net.IPv4len {
		ret.ip[10] = 0xff
		ret.ip[11] = 0xff
	}
	copy(ret.ip[net.IPv6len-len(b):net.IPv6len], b)

	ret.MakeHostInplace()
	return ret
}

// FromBigInt creates an IP from a net.IP
func (ip *IP) FromBigInt(i *big.Int) {
	b := i.Bytes()
	ip.ip = make(net.IP, net.IPv6len)
	if len(b) > net.IPv6len {
		return
	}
	if len(b) <= net.IPv4len {
		ip.ip[10] = 0xff
		ip.ip[11] = 0xff
	}
	copy(ip.ip[net.IPv6len-len(b):net.IPv6len], b)

	ip.MakeHostInplace()
}

// Parse creates an IP from a string
func Parse(s string) *IP {
	ip := &IP{}
	if s == "" {
		return Parse("0.0.0.0/0")
	}
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
	if i := strings.LastIndex(s, " "); i > -1 {
		address := string([]byte(s)[0:i])
		mask := string([]byte(s)[i+1 : len(s)])
		debugln(i, address, mask)
		ip.ip = net.ParseIP(address)
		if ip.ip == nil {
			debugf("ip parse error %s", address)
			return nil
		}
		m := net.ParseIP(mask)
		if m == nil {
			debugf("ip parse error %s", address)
			return nil
		}
		ip.mask, _ = net.IPMask(m.To4()).Size()
		if !MaskOk(ip.ip, ip.mask) {
			return nil
		}
		return ip
	}
	ip.ip = net.ParseIP(s)
	if ip.ip == nil {
		fmt.Println(ip.ip)
		return nil
	}
	ip.MakeHostInplace()
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

// MakeHost clones and sets the clone's mask to a host mask, then returns it
func (ip *IP) MakeHost() *IP {
	ret := ip.Clone()
	if v4 := ret.ip.To4(); v4 != nil {
		ret.mask = 32
	} else {
		ret.mask = 128
	}
	return ret
}

// MakeHostInplace sets the mask to a host mask
func (ip *IP) MakeHostInplace() {
	if v4 := ip.ip.To4(); v4 != nil {
		ip.mask = 32
	} else {
		ip.mask = 128
	}
}

// SetMask clones and sets the clone's mask to a host mask, then returns it
func (ip *IP) SetMask(m int) *IP {
	ret := ip.Clone()
	if MaskOk(ip.ip, m) {
		ret.mask = m
	}
	return ret
}

// SetMaskInplace sets the mask to a host mask
func (ip *IP) SetMaskInplace(m int) {
	if MaskOk(ip.ip, m) {
		ip.mask = m
	}
}

// MaskBits returns the Mask
func (ip *IP) MaskBits() int {
	return ip.mask
}

// IP returns the IP
func (ip *IP) IP() net.IP {
	return ip.ip
}

// IsV4 returns true if ip is an IPv4 address
func (ip *IP) IsV4() bool {
	return ip.ip.To4() != nil
}

// IsV6 returns true if ip is an IPv6 address
func (ip *IP) IsV6() bool {
	return len(ip.ip) == 16 && ip.ip.To4() == nil
}
