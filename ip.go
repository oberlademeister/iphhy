package iphhy

import (
	"net"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// I4 is a simple and easy to use type to handle an IP address or an IP network
// it is designed to be mutable
type I4 struct {
	ip       uint32
	maskBits int
}

func parseDottedQuad(dq string) (uint32, error) {
	var out uint32
	quads := strings.Split(dq, ".")
	if len(quads) != 4 {
		return 0, errors.New("parse error")
	}
	for i := 0; i < 4; i++ {
		val, err := strconv.Atoi(quads[i])
		if err != nil {
			return 0, errors.Wrap(err, "parse error")
		}
		if (val < 0) || (val > 255) {
			return 0, errors.New("parse error")
		}
		add := val * (1 << uint(24-8*i))
		if add > 0xFFFFFFFF {
			panic("something's wrong")
		}
		out += uint32(add)
	}
	return out, nil
}

// FromString sets the state from the string
func (i4 *I4) FromString(s string) error {
	i4.ip = 0
	i4.maskBits = 0
	// first, split the IP from the mask (separated by either a space or a /)
	slashSplit := strings.Split(s, "/")
	spaceSplit := strings.Split(s, " ")
	var dq string
	var err error
	var newIP uint32
	var newMask int

	switch {
	case len(slashSplit) == 2:
		dq = slashSplit[0]
		newMask, err = strconv.Atoi(slashSplit[1])
		if err != nil {
			return err
		}
	case len(spaceSplit) == 2:
		dq = spaceSplit[0]
		newMask = getMaskBitsFromString(spaceSplit[1])
	default:
		dq = s
		newMask = 32
	}
	if newMask < 0 || newMask > 32 {
		return errors.New("illegal mask")
	}
	newIP, err = parseDottedQuad(dq)
	if err != nil {
		return err
	}
	i4.ip = newIP
	i4.maskBits = newMask
	return nil
}

// NewI4 create a new I4 from a string
func NewI4(s string) (I4, error) {
	i := I4{}
	err := i.FromString(s)
	return i, err
}

// MustNewI4 creates a new I4 value and panics on error
func MustNewI4(s string) I4 {
	ai, err := NewI4(s)
	if err != nil {
		panic(err)
	}
	return ai
}

// IP returns the IP
func (i4 I4) IP() net.IP {
	return IntToIP(i4.ip)
}

// Mask returns the Mask
func (i4 I4) Mask() net.IPMask {
	return net.CIDRMask(i4.maskBits, 32)
}

// MaskBits returns the Mask
func (i4 I4) MaskBits() int {
	return i4.maskBits
}
