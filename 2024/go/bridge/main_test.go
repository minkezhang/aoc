package main

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
			data: []byte(`190: 10 19`),
			want: 190,
		},
		{
			name: "Example",
			data: []byte(`190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
`),
			want: 3749,
		},
		{
			name: "Trivial/Identity",
			data: []byte(`100: 1 1 1 1 1 1 4 1 1 1 1 25 1 1 1 1 1 1`),
			want: 100,
		},
	} {
		t.Run(c.name, func(t *testing.T) {
			if got := one(c.data); got != c.want {
				t.Errorf("one() = %v, want = %v", got, c.want)
			}
		})
	}
}

func TestCheck(t *testing.T) {
	for _, c := range []struct {
		name   string
		target int
		xs     []int
		want   bool
	}{
		{
			name:   "Trivial",
			target: 10,
			xs:     []int{10},
			want:   true,
		},
		{
			name:   "Trivial/Identity",
			target: 100,
			xs:     []int{1, 1, 1, 1, 1, 25, 1, 1, 1, 1, 4, 1, 1, 1, 1, 1},
			want:   true,
		},
		{
			name:   "Trivial/Multiply",
			target: 24,
			xs:     []int{6, 4},
			want:   true,
		},
		{
			name:   "Trivial/False",
			target: 10,
			xs:     []int{1},
			want:   false,
		},
		{
			name:   "Example/T=190",
			target: 190,
			xs:     []int{10, 19},
			want:   true,
		},
		{
			name:   "Example/T=3267",
			target: 3267,
			xs:     []int{81, 40, 27},
			want:   true,
		},
		{
			name:   "Example/T=83",
			target: 83,
			xs:     []int{17, 5},
			want:   false,
		},
	} {
		t.Run(c.name, func(t *testing.T) {
			if got := check(c.target, c.xs); got != c.want {
				t.Errorf("check() = %v, want = %v", got, c.want)
			}
		})
	}
}
