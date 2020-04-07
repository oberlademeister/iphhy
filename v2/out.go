package iphhy

import "fmt"

// String returns the string/CIDR mask notation
func (ip IP) String() string {
	return fmt.Sprintf("%s/%d", ip.ip.String(), ip.mask)
}
