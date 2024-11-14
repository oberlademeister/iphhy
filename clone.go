package iphhy

import "net"

// Clone returns a (deep) clone of the ip struct
func (ip *IP) Clone() *IP {
	ret := &IP{
		ip:   make(net.IP, len(ip.ip)),
		mask: ip.mask,
	}
	copy(ret.ip, ip.ip)
	return ret
}
