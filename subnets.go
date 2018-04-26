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
