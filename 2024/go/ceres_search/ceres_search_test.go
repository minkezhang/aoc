package ceres_search

import (
	"testing"
)

func TestOne(t *testing.T) {
	for _, c := range []struct {
		name string
		data []string
		want int
	}{
		{
			name: "Trivial",
			data: []string{
				"AAAA",
				"AAAA",
				"AAAA",
				"AAAA",
			},
			want: 0,
		},
		{
			name: "Horizontal",
			data: []string{
				"XMAS",
			},
			want: 1,
		},
		{
			name: "Vertical",
			data: []string{
				"X",
				"M",
				"A",
				"S",
			},
			want: 1,
		},
		{
			name: "Example",
			data: []string{
				"MMMSXXMASM",
				"MSAMXMSMSA",
				"AMXSXMAAMM",
				"MSAMASMSMX",
				"XMASAMXAMM",
				"XXAMMXXAMA",
				"SMSMSASXSS",
				"SAXAMASAAA",
				"MAMMMXMMMM",
				"MXMXAXMASX",
			},
			want: 18,
		},
	} {
		t.Run(c.name, func(t *testing.T) {
			if got := one(c.data); got != c.want {
				t.Errorf("one() = %v, want = %v", got, c.want)
			}
		})
	}
}
