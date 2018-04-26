package iphhy

import "testing"

func TestI4_Overlaps(t *testing.T) {
	tests := []struct {
		name string
		i1   I4
		i2   I4
		want bool
	}{
		{
			name: "disjunct",
			i1:   MustNewI4("192.168.1.0/24"),
			i2:   MustNewI4("192.168.2.0/24"),
			want: false,
		},
		{
			name: "same",
			i1:   MustNewI4("192.168.1.0/24"),
			i2:   MustNewI4("192.168.1.0/24"),
			want: true,
		},
		{
			name: "left overlap",
			i1:   MustNewI4("192.168.1.0/24"),
			i2:   MustNewI4("192.168.1.0/25"),
			want: true,
		},
		{
			name: "right overlap",
			i1:   MustNewI4("192.168.1.0/24"),
			i2:   MustNewI4("192.168.1.128/25"),
			want: true,
		},
		{
			name: "in between",
			i1:   MustNewI4("192.168.1.0/24"),
			i2:   MustNewI4("192.168.1.128/26"),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i1.Overlaps(tt.i2); got != tt.want {
				t.Errorf("I4.Overlaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
