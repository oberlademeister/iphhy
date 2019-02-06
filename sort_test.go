package iphhy

import (
	"reflect"
	"sort"
	"testing"
)

func TestI4Numerical(t *testing.T) {
	tests := []struct {
		name string
		in   []I4
		want []I4
	}{
		{
			name: "t1",
			in:   []I4{MustNewI4("10.0.0.0/8"), MustNewI4("10.0.0.0/24"), MustNewI4("11.0.0.0/8"), MustNewI4("9.0.0.0/8")},
			want: []I4{MustNewI4("9.0.0.0/8"), MustNewI4("10.0.0.0/8"), MustNewI4("10.0.0.0/24"), MustNewI4("11.0.0.0/8")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := I4Numerical(tt.in)
			sort.Sort(got)
			if !reflect.DeepEqual(tt.in, tt.want) {
				t.Errorf("I4NumericalSort = %v, want %v", got, tt.want)
			}
		})
	}
}
