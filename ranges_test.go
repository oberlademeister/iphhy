package iphhy

import (
	"reflect"
	"testing"
)

func TestI4_GetAll(t *testing.T) {
	type args struct {
		lowerOffset int
		upperOffset int
	}
	tests := []struct {
		name    string
		i       I4
		args    args
		want    []I4
		wantErr bool
	}{
		{
			name: "t1",
			i:    MustNewI4("192.168.1.0/29"),
			args: args{
				lowerOffset: 0,
				upperOffset: -1,
			},
			want: []I4{
				MustNewI4("192.168.1.0/29"),
				MustNewI4("192.168.1.1/29"),
				MustNewI4("192.168.1.2/29"),
				MustNewI4("192.168.1.3/29"),
				MustNewI4("192.168.1.4/29"),
				MustNewI4("192.168.1.5/29"),
				MustNewI4("192.168.1.6/29"),
				MustNewI4("192.168.1.7/29"),
			},
			wantErr: false,
		},
		{
			name: "boundswrong",
			i:    MustNewI4("192.168.1.0/29"),
			args: args{
				lowerOffset: 3,
				upperOffset: 2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "illegal lower offset",
			i:    MustNewI4("192.168.1.0/29"),
			args: args{
				lowerOffset: 30,
				upperOffset: 2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "illegal upper offset",
			i:    MustNewI4("192.168.1.0/29"),
			args: args{
				lowerOffset: 3,
				upperOffset: 30,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.i.GetAll(tt.args.lowerOffset, tt.args.upperOffset)
			if (err != nil) != tt.wantErr {
				t.Errorf("I4.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("I4.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
