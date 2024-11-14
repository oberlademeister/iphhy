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

func TestIP_SubnetOffset(t *testing.T) {
	type fields struct {
		ip   net.IP
		mask int
	}
	type args struct {
		offset int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *IP
		wantErr bool
	}{
		{
			name: "subnetoffset+0",
			fields: fields{
				ip:   net.ParseIP("192.168.0.1"),
				mask: 24,
			},
			args: args{
				offset: 0,
			},
			want:    Parse("192.168.0.0/24"),
			wantErr: false,
		},
		{
			name: "subnetoffset+2",
			fields: fields{
				ip:   net.ParseIP("192.168.0.1"),
				mask: 24,
			},
			args: args{
				offset: 2,
			},
			want:    Parse("192.168.0.2/24"),
			wantErr: false,
		},
		{
			name: "subnetoffset-1",
			fields: fields{
				ip:   net.ParseIP("192.168.0.1"),
				mask: 24,
			},
			args: args{
				offset: -1,
			},
			want:    Parse("192.168.0.255/24"),
			wantErr: false,
		},
		{
			name: "subnetoffset-2",
			fields: fields{
				ip:   net.ParseIP("192.168.0.1"),
				mask: 24,
			},
			args: args{
				offset: -2,
			},
			want:    Parse("192.168.0.254/24"),
			wantErr: false,
		},
		{
			name: "subnetoffset+256",
			fields: fields{
				ip:   net.ParseIP("192.168.0.1"),
				mask: 24,
			},
			args: args{
				offset: 256,
			},
			want:    &IP{ip: nil, mask: 0},
			wantErr: true,
		},
		{
			name: "subnetoffset-257",
			fields: fields{
				ip:   net.ParseIP("192.168.0.1"),
				mask: 24,
			},
			args: args{
				offset: -257,
			},
			want:    &IP{ip: nil, mask: 0},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ip := &IP{
				ip:   tt.fields.ip,
				mask: tt.fields.mask,
			}
			got, err := ip.SubnetOffset(tt.args.offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("IP.SubnetOffset() error = %v, want %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IP.SubnetOffset() = %+v, want %+v", got, tt.want)
			}

		})
	}
}
