package iphhy

import "fmt"

// String returns the IP in CIDR notation
func (i4 I4) String() string {
	return i4.IPString() + "/" + fmt.Sprint(i4.maskBits)
}

// CIDRString returns the IP in CIDR notation
func (i4 I4) CIDRString() string {
	return i4.IPString() + "/" + fmt.Sprint(i4.maskBits)
}

// DoubleDottedQuad returns the IP in double dotted quad 192.168.0.1 255.255.255.255
func (i4 I4) DoubleDottedQuad() string {
	return i4.IPString() + " " + v4MaskStrings[i4.maskBits]
}

// IPString returns the IP
func (i4 I4) IPString() string {
	var out string
	for i := 0; i < 4; i++ {
		b := (i4.ip >> uint32(24-8*i)) & 0xFF
		out += fmt.Sprintf("%d.", b)
	}
	return out[0 : len(out)-1]
}

// MaskString returns the mask in dq
func (i4 I4) MaskString() string {
	return v4MaskStrings[i4.maskBits]
}
