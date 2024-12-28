package guard_gallivant

import (
	"testing"
)

func TestOne(t *testing.T) {
	for _, c := range []struct {
		name string
		data string
		want int
	}{
		{
			name: "Trivial",
			data: `^`,
			want: 1,
		},
		{
			name: "Slice",
			data: `.
^`,
			want: 2,
		},
		{
			name: "Turn",
			data: `###
^..`,
			want: 3,
		},
		{
			name: "UTurn",
			data: `###
^#.`,
			want: 1,
		},
		{
			name: "Visited",
			data: `#..
..#
^..
.#.`,
			want: 4,
		},
		{
			name: "Example",
			data: `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`,
			want: 41,
		},
		{
			name: "Loop",
			data: `.#..
.^.#
#...
..#.`,
			want: -1,
		},
		{
			name: "Loop/NonStart",
			data: `#.#..
....#
^#...
...#.
`,
			want: -1,
		},
	} {
		t.Run(c.name, func(t *testing.T) {
			if got := one(c.data); got != c.want {
				t.Errorf("one() = %v, want = %v", got, c.want)
			}
		})
	}
}
