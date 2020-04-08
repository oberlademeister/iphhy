package iphhy

import (
	"net"
)

// Network returns a new *IP with the netmask applied
func (i *IP) Network() *IP {
	cl := i.Clone()
	cl.NetworkInPlace()
	return cl
}

// NetworkInPlace returns a new *IP with the netmask applied
func (i *IP) NetworkInPlace() {
	var bits int
	if i.IsV4() {
		bits = 32
	} else {
		bits = 128
	}
	mask := net.CIDRMask(i.mask, bits)
	i.ip = i.ip.Mask(mask)
}

// Broadcast returns a new *IP with the netmask applied
func (i *IP) Broadcast() *IP {
	cl := i.Clone()
	cl.BroadcastInPlace()
	return cl
}

// BroadcastInPlace returns a new *IP with the netmask applied
func (i *IP) BroadcastInPlace() {
	var address []byte
	var mask net.IPMask
	if i.IsV4() {
		address = i.ip.To4()
		mask = net.CIDRMask(i.mask, 32)
	} else {
		address = i.ip.To16()
		mask = net.CIDRMask(i.mask, 128)
	}
	l := len(address)
	out := make(net.IP, l)
	for i := range address {
		out[i] = address[i] | ^mask[i]
	}
	i.ip = out
}
