package iphhy

import (
	"math/big"
	"net"
)

// To4 returns the result of net.IP.To4()
func (ip *IP) To4() net.IP {
	return ip.ip.To4()
}

// To16 returns the result of net.IP.To4()
func (ip *IP) To16() net.IP {
	return ip.ip.To16()
}

// ToIP returns the ip field
func (ip *IP) ToIP() net.IP {
	return ip.ip
}

// BigInt returns a big int representing the value
func (ip *IP) BigInt() *big.Int {
	ret := &big.Int{}
	if i := ip.ip.To4(); i != nil {
		ret.SetBytes(i)
		return ret
	}
	ret.SetBytes(ip.ip)
	return ret
}
