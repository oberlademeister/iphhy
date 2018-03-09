package iph

import (
	"fmt"
	"net"
)

func invertMask(oldMask net.IPMask) net.IPMask {
	newMask := net.IPMask(make([]byte, 4))
	cpMask := make([]byte, 4)
	copy(cpMask, oldMask)
	for i := 0; i < 4; i++ {
		newMask[i] = ^cpMask[i]
	}
	return newMask
}

// BaseIP returns the network number
func (i I4) BaseIP() net.IP {
	m := net.CIDRMask(i.maskBits, 32)
	ip := IntToIP(i.ip)
	b := make([]byte, 4)
	//fmt.Printf("m: %08b ip: %08b b: %08b\n", m, ip, b)
	for i := 0; i < 4; i++ {
		b[i] = ip[i] & m[i]
	}
	//fmt.Printf("m: %08b ip: %08b b: %08b\n", m, ip, b)
	return b
}

// LastIP returns the last number in the network
func (i I4) LastIP() net.IP {
	m := invertMask(net.CIDRMask(i.maskBits, 32))
	ip := IntToIP(i.ip)
	b := make([]byte, 4)
	for i := 0; i < 4; i++ {
		b[i] = ip[i] | m[i]
	}
	return b
}

// Offset adds the offset to the network address
// offset 0 returns the network address
// offset 1 returns the first host address
// offset -1 returns the last address (i.e. the broadcast address)
// offset -2 returns the last host address
func (i *I4) Offset(offset int) (I4, error) {
	lower := IPToInt(i.BaseIP())
	upper := IPToInt(i.LastIP())
	var newIPInt uint32
	switch {
	case offset == 0:
		newIPInt = lower
	case offset < 0:
		newIPInt = upper - uint32(-1*offset-1)
	case offset > 0:
		newIPInt = lower + uint32(offset)
	}
	newIP := IntToIP(newIPInt)
	if (newIPInt < lower) || (newIPInt > upper) {
		return I4{}, fmt.Errorf("result not in subnet range %s+%d !in [%s]", newIP, offset, i.String())
	}
	return I4{newIPInt, i.maskBits}, nil
}

// OffsetString returns the IPv4 string for a given subnet string and a given offset
func OffsetString(subnet string, offset int) (string, error) {
	i4, err := NewI4(subnet)
	if err != nil {
		return "", err
	}
	s, err := i4.Offset(offset)
	if err != nil {
		return "", err
	}
	return s.IPString(), nil
}

// MustOffsetString same as OffsetString but panics if called with wrong parameter
func MustOffsetString(subnet string, offset int) string {
	s, err := OffsetString(subnet, offset)
	if err != nil {
		panic(err)
	}
	return s
}
