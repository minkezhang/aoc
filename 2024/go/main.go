package main

import (
	"github.com/minkezhang/aoc/2024/base"
	"github.com/minkezhang/aoc/2024/bridge_repair"
	"github.com/minkezhang/aoc/2024/ceres_search"
	"github.com/minkezhang/aoc/2024/disk_fragmenter"
	"github.com/minkezhang/aoc/2024/guard_gallivant"
	"github.com/minkezhang/aoc/2024/historian_hysteria"
	"github.com/minkezhang/aoc/2024/mull_it_over"
	"github.com/minkezhang/aoc/2024/print_queue"
	"github.com/minkezhang/aoc/2024/red_nosed_reports"
	"github.com/minkezhang/aoc/2024/resonant_collinearity"
)

var (
	ps = []base.P[int, int]{
		historian_hysteria.P{},
		red_nosed_reports.P{},
		mull_it_over.P{},
		ceres_search.P{},
		print_queue.P{},
		guard_gallivant.P{},
		bridge_repair.P{},
		resonant_collinearity.P{},
		disk_fragmenter.P{},
	}
)

func main() {
	for _, p := range ps {
		(&base.T[int, int]{P: p}).Print()
	}
}
