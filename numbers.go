package iph

import (
	"encoding/binary"
	"net"
)

// IPToInt converts an IPv4 address to a uint32
func IPToInt(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

// IntToIP converts a uint32 to an IPv4 address
func IntToIP(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}
