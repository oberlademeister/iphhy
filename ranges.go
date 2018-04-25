package iphhy

import "fmt"

// GetAll returns a slice of all I4 numbers, including first and last
func (i I4) GetAll(lowerOffset, upperOffset int) ([]I4, error) {
	lower, err := i.Offset(lowerOffset)
	if err != nil {
		return nil, fmt.Errorf("lowerOffset")
	}
	upper, err := i.Offset(upperOffset)
	if err != nil {
		return nil, fmt.Errorf("upperOffset")
	}
	if lower.Number() > upper.Number() {
		return nil, fmt.Errorf("lowerOffset has to result in lower ip than upperOffset")
	}
	var ret []I4
	for num := lower.Number(); num <= upper.Number(); num++ {
		ret = append(ret, I4{num, i.maskBits})
	}
	return ret, nil
}
