package mull_it_over

import (
	"testing"
)

func TestTwo(t *testing.T) {
	configs := []struct {
		name string
		data string
		want int
	}{
		{
			name: "Example",
			data: "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			want: 48,
		},
		{
			name: "AllDo",
			data: "mul(2,4)",
			want: 8,
		},
		{
			name: "AllDont",
			data: "don't()mul(2,4)",
			want: 0,
		},
		{
			name: "EndDont",
			data: "mul(2,4)don't()mul(9,9)",
			want: 8,
		},
		{
			name: "Alternate",
			data: "mul(2,4)don't()mul(9,9)do()mul(2,3)",
			want: 14,
		},
	}
	for _, c := range configs {
		t.Run(c.name, func(t *testing.T) {
			if got := two(c.data); got != c.want {
				t.Errorf("two() = %v, want = %v", got, c.want)
			}
		})
	}
}
