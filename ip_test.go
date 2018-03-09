package iph

import (
	"net"
	"reflect"
	"testing"
)

func TestNewI4(t *testing.T) {
	tests := []struct {
		Input      string
		WantedIP   uint32
		WantedMask int
		ErrIsNil   bool
	}{
		{"192.168.0.1", 3232235521, 32, true},
		{"192.168.0.1.1", 0, 0, false},
		{"192.XX.0.1", 0, 0, false},
		{"192.168.0.1/32", 3232235521, 32, true},
		{"192.168.0.1/23", 3232235521, 23, true},
		{"192.168.0.1/XXX", 0, 0, false},
		{"192.168.0.1 255.255.255.128", 3232235521, 25, true},
		{"322.168.0.1 255.255.255.128", 0, 0, false},
		{"192.168.0.1 255.255.255.131", 0, 0, false},
		{"192.168.0.1 255.255.255.131", 0, 0, false},
		{"192.168.0.1/33", 0, 0, false},
	}

	for i, tt := range tests {
		ai, err := NewI4(tt.Input)
		if (err == nil) != tt.ErrIsNil {
			t.Errorf("NewI4#%d: Failed! Wrong Error. wantedErrisNil: %t got: %t err: %v", i, tt.ErrIsNil, (err == nil), err)
		}
		if tt.WantedIP != ai.ip {
			t.Errorf("NewI4#%d: Failed! Wrong IP. wanted: %d got: %d", i, tt.WantedIP, ai.ip)
		}
		if tt.WantedMask != ai.maskBits {
			t.Errorf("NewI4#%d: Failed! Wrong Mask. wanted: %d, got: %d", i, tt.WantedMask, ai.maskBits)
		}
	}
}

func TestI4_DoubleDottedQuad(t *testing.T) {
	type fields struct {
		ip       uint32
		maskBits int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "10.0.0.1/24",
			fields: fields{
				ip:       167772161,
				maskBits: 24,
			},
			want: "10.0.0.1 255.255.255.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i4 := I4{
				ip:       tt.fields.ip,
				maskBits: tt.fields.maskBits,
			}
			if got := i4.DoubleDottedQuad(); got != tt.want {
				t.Errorf("I4.DoubleDottedQuad() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestI4_IP(t *testing.T) {
	type fields struct {
		ip       uint32
		maskBits int
	}
	tests := []struct {
		name   string
		fields fields
		want   net.IP
	}{
		{
			name: "10.0.0.1/24",
			fields: fields{
				ip:       167772161,
				maskBits: 24,
			},
			want: []byte{10, 0, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i4 := I4{
				ip:       tt.fields.ip,
				maskBits: tt.fields.maskBits,
			}
			if got := i4.IP(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("I4.IP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestI4_IPString(t *testing.T) {
	type fields struct {
		ip       uint32
		maskBits int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "10.0.0.1/24",
			fields: fields{
				ip:       167772161,
				maskBits: 24,
			},
			want: "10.0.0.1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i4 := I4{
				ip:       tt.fields.ip,
				maskBits: tt.fields.maskBits,
			}
			if got := i4.IPString(); got != tt.want {
				t.Errorf("I4.IPString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestI4_Mask(t *testing.T) {
	type fields struct {
		ip       uint32
		maskBits int
	}
	tests := []struct {
		name   string
		fields fields
		want   net.IPMask
	}{
		{
			name: "10.0.0.1/24",
			fields: fields{
				ip:       167772161,
				maskBits: 24,
			},
			want: net.CIDRMask(24, 32),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i4 := I4{
				ip:       tt.fields.ip,
				maskBits: tt.fields.maskBits,
			}
			if got := i4.Mask(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("I4.Mask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestI4_MaskBits(t *testing.T) {
	type fields struct {
		ip       uint32
		maskBits int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "10.0.0.1/24",
			fields: fields{
				ip:       167772161,
				maskBits: 24,
			},
			want: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i4 := I4{
				ip:       tt.fields.ip,
				maskBits: tt.fields.maskBits,
			}
			if got := i4.MaskBits(); got != tt.want {
				t.Errorf("I4.MaskBits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestI4_MaskString(t *testing.T) {
	type fields struct {
		ip       uint32
		maskBits int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "10.0.0.1/24",
			fields: fields{
				ip:       167772161,
				maskBits: 24,
			},
			want: "255.255.255.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i4 := I4{
				ip:       tt.fields.ip,
				maskBits: tt.fields.maskBits,
			}
			if got := i4.MaskString(); got != tt.want {
				t.Errorf("I4.MaskString() = %v, want %v", got, tt.want)
			}
		})
	}
}
