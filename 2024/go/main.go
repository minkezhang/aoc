package main

import (
	"github.com/minkezhang/aoc/2024/base"
	"github.com/minkezhang/aoc/2024/historian_hysteria"
	"github.com/minkezhang/aoc/2024/red_nosed_reports"
)

var (
	ps = []base.P[int, int]{
		historian_hysteria.P{},
		red_nosed_reports.P{},
	}
)

func main() {
	for _, p := range ps {
		(&base.T[int, int]{P: p}).Print()
	}
}
