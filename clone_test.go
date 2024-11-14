package iphhy

import (
	"reflect"
	"testing"
)

func TestIP_Clone(t *testing.T) {
	tests := []struct {
		name string
		ip   *IP
		want *IP
	}{
		{
			name: "t1",
			ip:   Parse("1.2.3.4/24"),
			want: Parse("1.2.3.4/24"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ip.Clone(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IP.Clone() = %v, want %v", got, tt.want)
			}
		})
	}
}
