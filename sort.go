package iphhy

// I4Numerical makes IP addresses sortable
type I4Numerical []I4

func (i4n I4Numerical) Len() int {
	return len(i4n)
}

func (i4n I4Numerical) Less(i, j int) bool {
	if i4n[i].ip == i4n[j].ip {
		return i4n[i].maskBits < i4n[j].maskBits
	}
	return i4n[i].ip < i4n[j].ip
}

func (i4n I4Numerical) Swap(i, j int) {
	i4n[i], i4n[j] = i4n[j], i4n[i]
}
