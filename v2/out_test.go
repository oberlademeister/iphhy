package iphhy

import (
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		input                          string
		wantedString                   string
		wantedIPString                 string
		wantedDottedQuadString         string
		wantedInvertedDottedQuadString string
		wantedMaskString               string
	}{
		{"192.168.21.18/24", "192.168.21.18/24", "192.168.21.18", "192.168.21.18 255.255.255.0", "192.168.21.18 0.0.0.255", "255.255.255.0"},
		{"2001:db8::/64", "2001:db8::/64", "2001:db8::", "2001:db8:: ffffffffffffffff0000000000000000", "2001:db8:: 0000000000000000ffffffffffffffff", "ffffffffffffffff0000000000000000"},
		{"2001:DB8::/64", "2001:db8::/64", "2001:db8::", "2001:db8:: ffffffffffffffff0000000000000000", "2001:db8:: 0000000000000000ffffffffffffffff", "ffffffffffffffff0000000000000000"},
	}

	for i, tt := range tests {
		aip := Parse(tt.input)
		s := aip.String()
		is := aip.IPString()
		dqs := aip.DoubleDottedQuad()
		dqivs := aip.DoubleDottedQuadInvertedMask()
		ms := aip.MaskString()
		switch {
		case s != tt.wantedString:
			t.Errorf("TestString#%d: tt: string / String failed: wanted %q got %q", i, tt.wantedString, s)
		case is != tt.wantedIPString:
			t.Errorf("TestString#%d: tt: string / IPString failed: wanted %q got %q", i, tt.wantedIPString, is)
		case dqs != tt.wantedDottedQuadString:
			t.Errorf("TestString#%d: tt: string / DDQ failed: wanted %q got %q", i, tt.wantedDottedQuadString, dqs)
		case dqivs != tt.wantedInvertedDottedQuadString:
			t.Errorf("TestString#%d: tt: string / DDQIV failed: wanted %q got %q", i, tt.wantedInvertedDottedQuadString, dqivs)
		case ms != tt.wantedMaskString:
			t.Errorf("TestString#%d: tt: string / maskstring failed: %q got %q", i, tt.wantedMaskString, ms)
		}
	}
}
