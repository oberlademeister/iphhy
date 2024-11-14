package iphhy

import (
	"fmt"
	"math/big"
)

// Offset returns a fresh IP struct set off by int
func (ip *IP) Offset(offset int64) *IP {
	i := ip.BigInt()
	i.Add(i, big.NewInt(offset))
	ret := NewFromBigInt(i)
	ret.SetMaskInplace(ip.mask)
	return ret
}

// SubnetOffset returns a fresh IP struct set off by int
func (ip *IP) SubnetOffset(offset int64) (*IP, error) {
	lower := ip.Network().BigInt()
	upper := ip.Broadcast().BigInt()
	var newIPInt = big.NewInt(0)
	switch {
	case offset == 0:
		newIPInt = lower
	case offset < 0:
		o := big.NewInt(-1*offset - 1)
		newIPInt = newIPInt.Sub(upper, o)
	case offset > 0:
		o := big.NewInt(offset)
		newIPInt = newIPInt.Add(lower, o)
	}
	// newIPInt < lower || newIPInt > upper
	newIP := NewFromBigInt(newIPInt)
	newIP.SetMaskInplace(ip.mask)
	if (newIPInt.Cmp(lower) == -1) || (newIPInt.Cmp(upper) == 1) {
		return &IP{}, fmt.Errorf("result not in subnet range %s+%d !in [%s]", newIP, offset, ip.String())
	}
	return newIP, nil
}

// MustSubnetOffset sames as SubnetOffset but panics on error
func (ip *IP) MustSubnetOffset(offset int64) *IP {
	i2, err := ip.SubnetOffset(offset)
	if err != nil {
		panic(err)
	}
	return i2
}
