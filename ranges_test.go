package iph

import (
	"net"
	"reflect"
	"testing"
)

func TestI4_GetAllIPs(t *testing.T) {
	type fields struct {
		ip       uint32
		maskBits int
	}
	tests := []struct {
		name   string
		fields fields
		want   []net.IP
	}{
		{
			name: "192.168.0.0/24",
			fields: fields{
				ip:       3232235520,
				maskBits: 29,
			},
			want: []net.IP{
				[]byte{192, 168, 0, 0},
				[]byte{192, 168, 0, 1},
				[]byte{192, 168, 0, 2},
				[]byte{192, 168, 0, 3},
				[]byte{192, 168, 0, 4},
				[]byte{192, 168, 0, 5},
				[]byte{192, 168, 0, 6},
				[]byte{192, 168, 0, 7},
			},
		},
	}
	for _, tt := range tests {
		i := I4{
			ip:       tt.fields.ip,
			maskBits: tt.fields.maskBits,
		}
		if got := i.GetAllIPs(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. I4.GetAllIPs() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestI4_GetAllIPStrings(t *testing.T) {
	type fields struct {
		ip       uint32
		maskBits int
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "192.168.0.0/24",
			fields: fields{
				ip:       3232235520,
				maskBits: 29,
			},
			want: []string{
				"192.168.0.0",
				"192.168.0.1",
				"192.168.0.2",
				"192.168.0.3",
				"192.168.0.4",
				"192.168.0.5",
				"192.168.0.6",
				"192.168.0.7",
			},
		},
	}
	for _, tt := range tests {
		i := I4{
			ip:       tt.fields.ip,
			maskBits: tt.fields.maskBits,
		}
		if got := i.GetAllIPStrings(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. I4.GetAllIPStrings() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
