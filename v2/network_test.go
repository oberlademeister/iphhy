package iphhy

import (
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
