package iphhy

import (
	"reflect"
	"testing"
)

func TestI4_Overlaps(t *testing.T) {
	tests := []struct {
		name string
		i1   I4
		i2   I4
		want bool
	}{
		{
			name: "disjunct",
			i1:   MustNewI4("192.168.1.0/24"),
			i2:   MustNewI4("192.168.2.0/24"),
			want: false,
		},
		{
			name: "same",
			i1:   MustNewI4("192.168.1.0/24"),
			i2:   MustNewI4("192.168.1.0/24"),
			want: true,
		},
		{
			name: "left overlap",
			i1:   MustNewI4("192.168.1.0/24"),
			i2:   MustNewI4("192.168.1.0/25"),
			want: true,
		},
		{
			name: "right overlap",
			i1:   MustNewI4("192.168.1.0/24"),
			i2:   MustNewI4("192.168.1.128/25"),
			want: true,
		},
		{
			name: "in between",
			i1:   MustNewI4("192.168.1.0/24"),
			i2:   MustNewI4("192.168.1.128/26"),
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

func TestI4_NumIPs(t *testing.T) {
	tests := []struct {
		name string
		i    I4
		want int
	}{
		{"sn32", MustNewI4("192.168.0.1/32"), 1},
		{"sn31", MustNewI4("192.168.0.1/31"), 2},
		{"sn30", MustNewI4("192.168.0.1/30"), 4},
		{"sn24", MustNewI4("192.168.0.1/24"), 256},
		{"sn16", MustNewI4("192.168.0.1/16"), 65536},
		{"sn0", MustNewI4("192.168.0.1/0"), 4294967296},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.NumIPs(); got != tt.want {
				t.Errorf("I4.NumIPs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestI4_NumHosts(t *testing.T) {
	tests := []struct {
		name string
		i    I4
		want int
	}{
		{"sn32", MustNewI4("192.168.0.1/32"), 1},
		{"sn31", MustNewI4("192.168.0.1/31"), 2},
		{"sn30", MustNewI4("192.168.0.1/30"), 2},
		{"sn24", MustNewI4("192.168.0.1/24"), 254},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.NumHosts(); got != tt.want {
				t.Errorf("I4.NumHosts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestI4_SubnetOffset(t *testing.T) {
	tests := []struct {
		name string
		i    I4
		o    int
		want I4
	}{
		{"t1", MustNewI4("192.168.0.1/24"), -1, MustNewI4("192.167.255.1/24")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.SubnetOffset(tt.o); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("I4.SubnetOffset() = %v, want %v", got, tt.want)
			}
		})
	}
}
