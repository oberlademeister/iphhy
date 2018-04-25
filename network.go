package iph

// Rebase replaces the IP address with the base (network) address
func (i *I4) Rebase() {
	i.ip = IPToInt(i.BaseIP())
}

// Network returns a new IP which is rebased
func (i I4) Network() I4 {
	return I4{IPToInt(i.BaseIP()), i.maskBits}
}
