package iphhy

import (
	"reflect"
	"sort"
	"testing"
)

func TestCompare(t *testing.T) {
	type args struct {
		ipA *IP
		ipB *IP
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "t1",
			args: args{
				ipA: Parse("1.2.3.4"),
				ipB: Parse("1.2.3.4"),
			},
			want: 0,
		},
		{
			name: "t2",
			args: args{
				ipA: Parse("1.2.3.4/24"),
				ipB: Parse("1.2.3.4/32"),
			},
			want: 1,
		},
		{
			name: "t2inv",
			args: args{
				ipA: Parse("1.2.3.4/32"),
				ipB: Parse("1.2.3.4/24"),
			},
			want: -1,
		},
		{
			name: "t3",
			args: args{
				ipA: Parse("1.2.3.4/24"),
				ipB: Parse("1.2.3.5/32"),
			},
			want: -1,
		},
		{
			name: "t4nilA",
			args: args{
				ipA: nil,
				ipB: Parse("1.2.3.5/32"),
			},
			want: -1,
		},
		{
			name: "t4nilB",
			args: args{
				ipA: Parse("1.2.3.5/32"),
				ipB: nil,
			},
			want: 1,
		},
		{
			name: "t4nilBoth",
			args: args{
				ipA: nil,
				ipB: nil,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compare(tt.args.ipA, tt.args.ipB); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSort(t *testing.T) {
	unsorted := IPList{
		Parse("1.2.3.5"),
		Parse("1.2.3.4/24"),
		Parse("1.2.3.4"),
	}

	sorted := IPList{
		Parse("1.2.3.4"),
		Parse("1.2.3.4/24"),
		Parse("1.2.3.5"),
	}
	sort.Sort(unsorted)
	if !reflect.DeepEqual(unsorted, sorted) {
		t.Errorf("sorting failed, want %v got %v", sorted, unsorted)
	}
}
