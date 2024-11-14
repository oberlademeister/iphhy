package iphhy

import (
	"math/big"
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

// Overlaps determines if the two subnets share nodes
func (ip *IP) Overlaps(ip2 *IP) bool {
	l1 := ip.Network().BigInt()
	u1 := ip.Broadcast().BigInt()
	l2 := ip2.Network().BigInt()
	u2 := ip2.Broadcast().BigInt()
	if l1.Cmp(u2) > 0 || l2.Cmp(u1) > 0 {
		return false
	}
	return true
}

// NumIPs gives the number of IPs in the subnet
func (ip IP) NumIPs() *big.Int {
	ret := big.NewInt(1)
	m := big.NewInt(2)
	l := 0
	if ip.IsV4() {
		l = 32 - ip.mask
	}
	if ip.IsV6() {
		l = 128 - ip.mask
	}
	for i := 0; i < l; i++ {
		ret.Mul(ret, m)
	}
	return ret
}
