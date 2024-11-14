package iphhy

import (
	"fmt"
	"net"
)

// String returns the string/CIDR mask notation
func (ip IP) String() string {
	if len(ip.ip) > 0 {
		return fmt.Sprintf("%s/%d", ip.ip.String(), ip.mask)
	}
	return fmt.Sprintf("0.0.0.0/%d", ip.mask)
}

// IPString returns the IP
func (ip IP) IPString() string {
	if len(ip.ip) > 0 {
		return fmt.Sprintf("%s", ip.ip.String())
	}
	return "0.0.0.0"
}

// DoubleDottedQuad returns a double dotted quad string
// if the address is v6, the IP address plus the hexadecimal form
// of the mask is printed
func (ip IP) DoubleDottedQuad() string {
	var ret string
	ret += ip.ip.String() + " " + ip.MaskString()
	return ret
}

// DoubleDottedQuadInvertedMask returns a double dotted quad string
// if the address is v6, the IP address plus the hexadecimal form
// of the mask is printed
func (ip IP) DoubleDottedQuadInvertedMask() string {
	var ret string
	ret += ip.ip.String() + " " + ip.InvertedMaskString()
	return ret
}

// MaskString returns the mask in dq
func (ip IP) MaskString() string {
	var ret string
	if ip.IsV4() {
		mask := net.CIDRMask(ip.mask, 32)
		ret = net.IP(mask).String()
		return ret
	}
	mask := net.CIDRMask(ip.mask, 128)
	ret = mask.String()
	return ret
}

// InvertedMaskString returns the mask in dq
func (ip IP) InvertedMaskString() string {
	var ret string
	if ip.IsV4() {
		mask := InvertIPMask(net.CIDRMask(ip.mask, 32))
		ret = net.IP(mask).String()
		return ret
	}
	mask := InvertIPMask(net.CIDRMask(ip.mask, 128))
	ret = mask.String()
	return ret
}

// InvertIPMask inverts a net.IPMask
func InvertIPMask(in net.IPMask) net.IPMask {
	out := make(net.IPMask, len(in))
	for i, b := range in {
		out[i] = ^b
	}
	return out
}
