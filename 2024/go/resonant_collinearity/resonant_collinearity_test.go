package resonant_collinearity

import (
	"testing"
)

func TestOne(t *testing.T) {
	for _, c := range []struct {
		name string
		data []byte
		want int
	}{
		{
			name: "Simple",
			data: []byte(`.`),
			want: 0,
		},
		{
			name: "Simple/Horizontal",
			data: []byte(`a.a..`),
			want: 1,
		},
		{
			name: "Simple/Horizontal/Reverse",
			data: []byte(`..a.a`),
			want: 1,
		},
		{
			name: "Simple/Horizontal/Overlap",
			data: []byte(`a.a.a`),
			want: 2,
		},
		{
			name: "DifferentFreq",
			data: []byte(`A.a..`),
			want: 0,
		},
		{
			name: "Diagonal",
			data: []byte(`.....
.....
.aA..
..aA.
.....
`),
			want: 4,
		},
		{
			name: "Example",
			data: []byte(`............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............
`),
			want: 14,
		},
	} {
		t.Run(c.name, func(t *testing.T) {
			if got := New(c.data).One(); got != c.want {
				t.Errorf("One() = %v, want = %v", got, c.want)
			}
		})
	}
}

func TestTwo(t *testing.T) {
	for _, c := range []struct {
		name string
		data []byte
		want int
	}{
		{
			name: "Simple",
			data: []byte(`.`),
			want: 0,
		},
		{
			name: "Simple/Horizontal",
			data: []byte(`a.a....`),
			want: 4,
		},
		{
			name: "Example",
			data: []byte(`............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............
`),
			want: 34,
		},
	} {
		t.Run(c.name, func(t *testing.T) {
			if got := New(c.data).Two(); got != c.want {
				t.Errorf("Two() = %v, want = %v", got, c.want)
			}
		})
	}
}
