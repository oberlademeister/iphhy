package iphhy

import "bytes"

// Compare is useful for sorting
func Compare(ipA, ipB *IP) int {
	if ipA == nil {
		if ipB == nil {
			return 0
		}
		return -1
	}
	if ipB == nil {
		return 1
	}
	c := bytes.Compare(ipA.ip, ipB.ip)
	if c == 0 {
		switch {
		case ipA.mask == ipB.mask:
			return 0
		case ipA.mask > ipB.mask:
			return -1
		case ipA.mask < ipB.mask:
			return 1
		}
	}
	return c
}

// IPList is a list of IPs
type IPList []*IP

// Len is used for sorting
func (ipl IPList) Len() int { return len(ipl) }

// Swap is used for sorting
func (ipl IPList) Swap(i, j int) { ipl[i], ipl[j] = ipl[j], ipl[i] }

// Less is used for sorting
func (ipl IPList) Less(i, j int) bool {
	return Compare(ipl[i], ipl[j]) == -1
}
