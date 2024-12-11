package main

import (
	"testing"
)

func TestOne(t *testing.T) {
	for _, c := range []struct {
		name  string
		edges [][]int
		seqs  [][]int
		want  int
	}{
		{
			name:  "Trivial",
			edges: [][]int{},
			seqs:  [][]int{[]int{100}, []int{1}, []int{24}},
			want:  125,
		},
		{
			name:  "Simple/Invalid",
			edges: [][]int{[]int{1, 2}},
			seqs:  [][]int{[]int{2, 1, 100}, []int{0, 2, 1}},
			want:  0,
		},
		{
			name:  "Simple/Valid",
			edges: [][]int{[]int{1, 2}},
			seqs:  [][]int{[]int{1, 2, 3}, []int{1}, []int{2}, []int{100, 1, 2, 101, 0}},
			want:  7,
		},
		{
			name: "Example",
			edges: [][]int{
				[]int{47, 53},
				[]int{97, 13},
				[]int{97, 61},
				[]int{97, 47},
				[]int{75, 29},
				[]int{61, 13},
				[]int{75, 53},
				[]int{29, 13},
				[]int{97, 29},
				[]int{53, 29},
				[]int{61, 53},
				[]int{97, 53},
				[]int{61, 29},
				[]int{47, 13},
				[]int{75, 47},
				[]int{97, 75},
				[]int{47, 61},
				[]int{75, 61},
				[]int{47, 29},
				[]int{75, 13},
				[]int{53, 13},
			},
			seqs: [][]int{
				[]int{75, 47, 61, 53, 29},
				[]int{97, 61, 53, 29, 13},
				[]int{75, 29, 13},
				[]int{75, 97, 47, 61, 53},
				[]int{61, 13, 29},
				[]int{97, 13, 75, 29, 47},
			},
			want: 143,
		},
	} {
		t.Run(c.name, func(t *testing.T) {
			if got := one(c.edges, c.seqs); got != c.want {
				t.Errorf("one() = %v, want = %v", got, c.want)
			}
		})
	}
}
