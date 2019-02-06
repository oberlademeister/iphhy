package iphhy

import (
	"net"
	"reflect"
	"testing"
)

func TestInvertIPMask(t *testing.T) {
	type args struct {
		in net.IPMask
	}
	tests := []struct {
		name string
		args args
		want net.IPMask
	}{
		{
			name: "t1",
			args: args{
				in: net.CIDRMask(32, 32),
			},
			want: net.CIDRMask(0, 32),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InvertIPMask(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InvertIPMask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIPMask_Invert(t *testing.T) {
	tests := []struct {
		name string
		m    IPMask
		want IPMask
	}{
		{
			name: "t1",
			m:    CIDRMask(16, 32),
			want: IPv4Mask(0, 0, 255, 255),
		},
		{
			name: "t2",
			m:    CIDRMask(17, 32),
			want: IPv4Mask(0, 0, 127, 255),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Invert(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IPMask.Invert() = %v, want %v", got, tt.want)
			}
		})
	}
}
