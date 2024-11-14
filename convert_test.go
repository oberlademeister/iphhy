package iphhy

import (
	"math/big"
	"reflect"
	"testing"
)

func TestIP_BigInt(t *testing.T) {
	tests := []struct {
		name string
		ip   *IP
		want *big.Int
	}{
		{
			name: "1.2.3.4",
			ip:   Parse("1.2.3.4"),
			want: big.NewInt(16909060),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ip.BigInt(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IP.BigInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
