package iph

import "net"

// GetAllIPs returns a slice with all IPs in the network
func (i I4) GetAllIPs() []net.IP {
	var out []net.IP
	lower := IPToInt(i.BaseIP())
	upper := IPToInt(i.LastIP())
	for i := lower; i <= upper; i++ {
		out = append(out, IntToIP(i))
	}
	return out
}

// GetAllIPStrings returns a slice with all IPs in the network
// withMask determines if the mask is added in CIDR notation
func (i I4) GetAllIPStrings() []string {
	var out []string
	lower := IPToInt(i.BaseIP())
	upper := IPToInt(i.LastIP())
	for i := lower; i <= upper; i++ {
		out = append(out, IntToIP(i).String())
	}
	return out
}
