package iphhy

import (
	"math/big"
	"testing"
)

func TestIP_Network(t *testing.T) {
	tests := []struct {
		name string
		ip   *IP
		want *IP
	}{
		{
			name: "network-t1",
			ip:   Parse("1.2.3.4/24"),
			want: Parse("1.2.3.0/24"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ip.Network(); Compare(got, tt.want) != 0 {
				t.Errorf("IP.Network() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIP_Broadcast(t *testing.T) {
	tests := []struct {
		name string
		ip   *IP
		want *IP
	}{
		{
			name: "broadcast-t1",
			ip:   Parse("1.2.3.4/24"),
			want: Parse("1.2.3.255/24"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ip.Broadcast(); Compare(got, tt.want) != 0 {
				t.Errorf("IP.Broadcast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIP_Overlaps(t *testing.T) {
	tests := []struct {
		name string
		i1   *IP
		i2   *IP
		want bool
	}{
		{
			name: "disjunct",
			i1:   Parse("192.168.1.0/24"),
			i2:   Parse("192.168.2.0/24"),
			want: false,
		},
		{
			name: "same",
			i1:   Parse("192.168.1.0/24"),
			i2:   Parse("192.168.1.0/24"),
			want: true,
		},
		{
			name: "left overlap",
			i1:   Parse("192.168.1.0/24"),
			i2:   Parse("192.168.1.0/25"),
			want: true,
		},
		{
			name: "right overlap",
			i1:   Parse("192.168.1.0/24"),
			i2:   Parse("192.168.1.128/25"),
			want: true,
		},
		{
			name: "in between",
			i1:   Parse("192.168.1.0/24"),
			i2:   Parse("192.168.1.128/26"),
			want: true,
		},
		{
			name: "disjunct",
			i1:   Parse("2001:db8::/32"),
			i2:   Parse("2001:db9::/32"),
			want: false,
		},
		{
			name: "same",
			i1:   Parse("2001:db8::/32"),
			i2:   Parse("2001:db8::/32"),
			want: true,
		},
		{
			name: "left overlap",
			i1:   Parse("2001:db8::/32"),
			i2:   Parse("2001:db8::/33"),
			want: true,
		},
		{
			name: "right overlap",
			i1:   Parse("2001:db8::/32"),
			i2:   Parse("2001:db8:8000::/33"),
			want: true,
		},
		{
			name: "in between",
			i1:   Parse("2001:db8::/32"),
			i2:   Parse("2001:db8:8000::/34"),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i1.Overlaps(tt.i2); got != tt.want {
				t.Errorf("I4.Overlaps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIP_NumIPs(t *testing.T) {
	verylarge := new(big.Int)
	verylarge.SetString("5192296858534827628530496329220096", 10)
	tests := []struct {
		name string
		i    *IP
		want *big.Int
	}{
		{"sn32", Parse("192.168.0.1/32"), big.NewInt(1)},
		{"sn31", Parse("192.168.0.1/31"), big.NewInt(2)},
		{"sn30", Parse("192.168.0.1/30"), big.NewInt(4)},
		{"sn24", Parse("192.168.0.1/24"), big.NewInt(256)},
		{"sn16", Parse("192.168.0.1/16"), big.NewInt(65536)},
		{"sn0", Parse("192.168.0.1/0"), big.NewInt(4294967296)},
		{"snv616", Parse("abcd:ef01::/16"), verylarge},
		{"snv6128", Parse("abcd:ef01::1/128"), big.NewInt(1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.NumIPs(); got.Cmp(tt.want) != 0 {
				t.Errorf("IP.NumIPs() = %v, want %v", got, tt.want)
			}
		})
	}
}
