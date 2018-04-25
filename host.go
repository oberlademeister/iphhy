package iphhy

// ToHost returns a new I4 with a /32 mask
func (i I4) ToHost() I4 {
	return I4{i.ip, 32}
}
