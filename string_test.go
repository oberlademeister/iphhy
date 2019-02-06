package iphhy

import (
	"testing"
)

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

func TestI4_DoubleDottedQuadInvertedMask(t *testing.T) {
	tests := []struct {
		name string
		in   I4
		want string
	}{
		{
			name: "t1",
			in:   MustNewI4("192.168.0.1 255.255.255.0"),
			want: "192.168.0.1 0.0.0.255",
		},
		{
			name: "t2",
			in:   MustNewI4("192.168.0.1 255.255.192.0"),
			want: "192.168.0.1 0.0.63.255",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.in.DoubleDottedQuadInvertedMask(); got != tt.want {
				t.Errorf("I4.DoubleDottedQuadInvertedMask() = %v, want %v", got, tt.want)
			}
		})
	}
}
