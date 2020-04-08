package iphhy

import "math/big"

// Offset returns a fresh IP struct set off by int
func (ip *IP) Offset(offset int64) *IP {
	i := ip.BigInt()
	i.Add(i, big.NewInt(offset))
	ret := NewFromBigInt(i)
	ret.SetMaskInplace(ip.mask)
	return ret
}
