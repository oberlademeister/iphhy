package iphhy

import "net"

// InvertIPMask inverts a net.IPMask
func InvertIPMask(in net.IPMask) net.IPMask {
	out := make(net.IPMask, len(in))
	for i, b := range in {
		out[i] = ^b
	}
	return out
}

// IPMask is like the net.IPMask type
type IPMask []byte

// CIDRMask points to net.CIDRMask
func CIDRMask(ones, bits int) IPMask {
	return IPMask(net.CIDRMask(ones, bits))
}

// IPv4Mask points to net.IPv4Mask
func IPv4Mask(a, b, c, d byte) IPMask {
	return IPMask(net.IPv4Mask(a, b, c, d))
}

// Size points to net.IPMask.Size
func (m IPMask) Size() (ones, bits int) {
	return net.IPMask(m).Size()
}

// String points to net.IPMask.String
func (m IPMask) String() string {
	return net.IPMask(m).String()
}

// Invert inverts the Mask and returns itself
func (m IPMask) Invert() IPMask {
	for i := range m {
		m[i] = ^m[i]
	}
	return m
}
