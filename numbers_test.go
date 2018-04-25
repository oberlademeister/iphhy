package iphhy

import (
	"net"
	"reflect"
	"testing"
)

func TestIPToInt(t *testing.T) {
	type args struct {
		ip net.IP
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "0",
			args: args{
				ip: []byte{0, 0, 0, 0},
			},
			want: 0,
		},
		{
			name: "10.0.0.0",
			args: args{
				ip: []byte{10, 0, 0, 0},
			},
			want: 167772160,
		},
	}
	for _, tt := range tests {
		if got := IPToInt(tt.args.ip); got != tt.want {
			t.Errorf("%q. IPToInt() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestIntToIP(t *testing.T) {
	type args struct {
		nn uint32
	}
	tests := []struct {
		name string
		args args
		want net.IP
	}{
		{
			name: "0",
			args: args{
				nn: 0,
			},
			want: []byte{0, 0, 0, 0},
		},
		{
			name: "10.0.0.0",
			args: args{
				nn: 167772160,
			},
			want: []byte{10, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		if got := IntToIP(tt.args.nn); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. IntToIP() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
