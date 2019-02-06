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
