package iphhy

// Host returns a new I4 with a /32 mask
func (i I4) Host() I4 {
	return I4{i.ip, 32}
}

// MakeHost sets the mask to 32
func (i *I4) MakeHost() {
	i.maskBits = 32
}
