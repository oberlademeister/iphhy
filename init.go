package iphhy

import "net"

var v4MasksFromDQ map[string]int
var v4MaskStrings [33]string

func init() {
	v4MasksFromDQ = make(map[string]int)
	for i := 0; i <= 32; i++ {
		m := net.IP(net.CIDRMask(i, 32)).String()
		v4MasksFromDQ[m] = i
		v4MaskStrings[i] = m
	}
}

func getMaskBitsFromString(s string) int {
	v, ok := v4MasksFromDQ[s]
	if !ok {
		return -1
	}
	return v
}

func getMaskInt(bits int) uint32 {
	return uint32(0xFFFFFFFF) << uint32(32-bits)
}
