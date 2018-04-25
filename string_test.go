package iphhy

import "testing"

func TestString(t *testing.T) {
	tests := []struct {
		input                  string
		wantedString           string
		wantedIPString         string
		wantedDottedQuadString string
		wantedMaskString       string
	}{
		{"192.168.21.18/24", "192.168.21.18/24", "192.168.21.18", "192.168.21.18 255.255.255.0", "255.255.255.0"},
	}

	for i, tt := range tests {
		aip, _ := NewI4(tt.input)
		scidr := aip.CIDRString()
		s := aip.String()
		is := aip.IPString()
		dqs := aip.DoubleDottedQuad()
		ms := aip.MaskString()
		switch {
		case scidr != tt.wantedString || s != tt.wantedString:
			t.Errorf("TestString#%d: tt: string / cidrstring failed: wanted %q got %q", i, tt.wantedString, s)
		case is != tt.wantedIPString:
			t.Errorf("TestString#%d: tt: string / IPString failed: wanted %q got %q", i, tt.wantedIPString, is)
		case dqs != tt.wantedDottedQuadString:
			t.Errorf("TestString#%d: tt: string / DDQ failed: wanted %q got %q", i, tt.wantedDottedQuadString, dqs)
		case ms != tt.wantedMaskString:
			t.Errorf("TestString#%d: tt: string / maskstring failed: %q got %q", i, tt.wantedMaskString, ms)
		}
	}
}
