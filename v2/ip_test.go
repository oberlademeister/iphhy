package iphhy

import (
	"net"
	"reflect"
	"testing"
)

func TestNewFromNetIP(t *testing.T) {
	type args struct {
		ip net.IP
	}
	tests := []struct {
		name string
		args args
		want *IP
	}{
		{
			name: "v4-1",
			args: args{
				ip: net.ParseIP("1.2.3.4"),
			},
			want: &IP{
				ip:   net.ParseIP("1.2.3.4"),
				mask: 32,
			},
		},
		{
			name: "v6-1",
			args: args{
				ip: net.ParseIP("2001:DB8::1"),
			},
			want: &IP{
				ip:   net.ParseIP("2001:DB8::1"),
				mask: 128,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromNetIP(tt.args.ip); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromNetIP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *IP
	}{
		{
			name: "empty",
			args: args{
				s: "",
			},
			want: &IP{ip: net.ParseIP("0.0.0.0"), mask: 0},
		},
		{
			name: "1.2.3.4/24",
			args: args{
				s: "1.2.3.4/24",
			},
			want: &IP{ip: net.ParseIP("1.2.3.4"), mask: 24},
		},
		{
			name: "1.2.3.4 255.255.255.0",
			args: args{
				s: "1.2.3.4 255.255.255.0",
			},
			want: &IP{ip: net.ParseIP("1.2.3.4"), mask: 24},
		},
		{
			name: "1.2.3.4/33",
			args: args{
				s: "1.2.3.4/33",
			},
			want: nil,
		},
		{
			name: "1.2.3.4/-1",
			args: args{
				s: "1.2.3.4/-1",
			},
			want: nil,
		},
		{
			name: "1.2.3.4/XX",
			args: args{
				s: "1.2.3.4/XX",
			},
			want: nil,
		},
		{
			name: "1.b.3.4",
			args: args{
				s: "1.b.3.4",
			},
			want: nil,
		},
		{
			name: "1.b.3.4/24",
			args: args{
				s: "1.b.3.4/24",
			},
			want: nil,
		},
		{
			name: "2001:DB8::1/96",
			args: args{
				s: "2001:DB8::1/96",
			},
			want: &IP{ip: net.ParseIP("2001:DB8::1"), mask: 96},
		},
		{
			name: "2001:DB8::1",
			args: args{
				s: "2001:DB8::1",
			},
			want: &IP{ip: net.ParseIP("2001:DB8::1"), mask: 128},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
