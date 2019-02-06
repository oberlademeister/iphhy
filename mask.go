package iphhy

import "net"

// InvertIPMask inverts an IP Mask
func InvertIPMask(in net.IPMask) net.IPMask {
	out := make(net.IPMask, len(in))
	for i, b := range in {
		out[i] = ^b
	}
	return out
}
