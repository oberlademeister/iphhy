package iphhy

// Overlaps determines if the two subnets share nodes
func (i I4) Overlaps(i2 I4) bool {
	l1 := i.Base().Number()
	u1 := i.Last().Number()
	l2 := i2.Base().Number()
	u2 := i2.Last().Number()
	if l1 > u2 || l2 > u1 {
		return false
	}
	return true
}

// NumIPs gives the number of IPs in the subnet
func (i I4) NumIPs() int {
	return 1 << uint32(32-i.maskBits)
}

// NumHosts gives the number of hosts
func (i I4) NumHosts() int {
	switch i.maskBits {
	case 32:
		return 1
	case 31:
		return 2
	case 30:
		return 2
	}
	return i.NumIPs() - 2
}

// SubnetOffset gives the +-i subnet with the same mask, i.e. the IP address that is shifted o times the number of hosts upwards or downwards
func (i I4) SubnetOffset(o int) I4 {
	return I4{ip: i.ip + uint32(o*i.NumIPs()), maskBits: i.maskBits}
}
