package iphhy

import (
	"fmt"
	"net"
)

// String returns the string/CIDR mask notation
func (ip IP) String() string {
	return fmt.Sprintf("%s/%d", ip.ip.String(), ip.mask)
}

// DoubleDottedQuad returns a double dotted quad string
// if the address is v6, the IP address plus the hexadecimal form
// of the mask is printed
func (ip IP) DoubleDottedQuad() string {
	var ret string
	if ip.IsV4() {
		ret += ip.ip.String() + " "
		mask := net.CIDRMask(ip.mask, 32)
		ret += net.IP(mask).String()
		return ret
	}
	ret += ip.ip.String() + " "
	mask := net.CIDRMask(ip.mask, 128)
	ret += mask.String()
	return ret
}
