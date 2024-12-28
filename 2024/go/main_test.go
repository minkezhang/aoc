package main

import (
	"testing"

	"github.com/minkezhang/aoc/2024/base"
)

func BenchmarkSuite(b *testing.B) {
	for _, p := range ps {
		t := &base.T[int, int]{P: p}
		t.ReadOrDie()
		b.Run(t.P.Name(), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				t.Solve()
			}
		})
	}
}
