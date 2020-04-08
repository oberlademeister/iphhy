package iphhy

import (
	"net"
	"reflect"
	"testing"
)

func TestIP_Offset(t *testing.T) {
	type fields struct {
		ip   net.IP
		mask int
	}
	type args struct {
		offset int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *IP
	}{
		{
			name: "offset+1",
			fields: fields{
				ip:   net.ParseIP("1.2.3.4"),
				mask: 24,
			},
			args: args{
				offset: 1,
			},
			want: Parse("1.2.3.5/24"),
		},
		{
			name: "offset-1",
			fields: fields{
				ip:   net.ParseIP("1.2.3.4"),
				mask: 24,
			},
			args: args{
				offset: -1,
			},
			want: Parse("1.2.3.3/24"),
		},
		{
			name: "offset+256",
			fields: fields{
				ip:   net.ParseIP("1.2.3.0"),
				mask: 24,
			},
			args: args{
				offset: 256,
			},
			want: Parse("1.2.4.0/24"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ip := &IP{
				ip:   tt.fields.ip,
				mask: tt.fields.mask,
			}
			if got := ip.Offset(tt.args.offset); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IP.Offset() = %v, want %v", got, tt.want)
			}
		})
	}
}
