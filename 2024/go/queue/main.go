package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	fn = "queue.txt"
)

func filter(edges [][]int, seqs [][]int, include bool) [][]int {
	rseqs := [][]int{}

	invalidPrev := map[int]map[int]bool{}
	invalidNext := map[int]map[int]bool{}
	for _, e := range edges {
		h, t := e[0], e[1]

		if _, ok := invalidPrev[h]; !ok {
			invalidPrev[h] = map[int]bool{}
		}
		invalidPrev[h][t] = true

		if _, ok := invalidNext[t]; !ok {
			invalidNext[t] = map[int]bool{}
		}
		invalidNext[t][h] = true
	}

	v := func(s []int) bool {
		cache := map[int]bool{}

		// forward check
		for i := 0; i < len(s); i++ {
			for p := range invalidPrev[s[i]] {
				if _, ok := cache[p]; ok { // invalid
					return false
				}
			}
			cache[s[i]] = true
		}

		cache = map[int]bool{}
		// reverse check
		for i := 0; i < len(s); i++ {
			for p := range invalidNext[s[len(s)-1-i]] {
				if _, ok := cache[p]; ok { // invalid
					return false
				}
			}
			cache[s[len(s)-1-i]] = true
		}

		return true
	}

	for _, s := range seqs {
		if v(s) == include {
			rseqs = append(rseqs, s)
		}
	}

	return rseqs
}

// two fixes a sequence to comply with the list of (H, T) pairs.
//
// The problem makes several assumptions which are not stated --
//
//  1. Only one valid sequence exists i.e.
//     a. All numbers in a sequence shows up as an edge head or tail
//     b. The set of relevant edges for each sequence must produce a totally
//     ordered sequence (e.g. 1|2; 3|4 is not a valid set of edges for the
//     sequence (1, 3, 2, 4).
//  2. There are no cycles in the list of edges
//
// Given these assumptions, we know that the last element T of a valid sequence
// will have no relevent edge E such that E = (X, T) for some X in the sequence.
//
// We then can remove this tail from the sequence and prune down the relevant
// edges, then iterate over the next "new" tail.
func two(edges [][]int, seqs [][]int) int {
	seqs = filter(edges, seqs, false)

	es := map[int]map[int]bool{}
	for _, e := range edges {
		h, t := e[0], e[1]

		if _, ok := es[h]; !ok {
			es[h] = map[int]bool{}
		}
		es[h][t] = true
	}

	acc := 0
	for _, seq := range seqs {
		sl := map[int]bool{}
		for _, i := range seq {
			sl[i] = true
		}

		// Both the head and tail must exist in the edge lookup.
		et := map[int]map[int]bool{}
		for _, s := range seq {
			et[s] = map[int]bool{}
		}

		for h, ts := range es {
			if _, ok := sl[h]; ok {
				for t := range ts {
					if _, ok := sl[t]; ok {
						et[h][t] = true
					}
				}
			}
		}

		rseq := []int{}
		for len(et) > 0 {
			var tail int
			for h, ts := range et {
				if len(ts) == 0 { // Found current tail.
					tail = h
					delete(et, h)
					break
				}
			}

			// First element in t is the last element in the
			// corrected sequence.
			rseq = append(rseq, tail)

			// Remove all references to the tail
			for _, ts := range et {
				if _, ok := ts[tail]; ok {
					delete(ts, tail)
				}
			}
		}
		// Median element is invariant under reverse.
		acc += rseq[(len(rseq)-1)/2]
	}

	return acc
}

func one(edges [][]int, seqs [][]int) int {
	seqs = filter(edges, seqs, true)

	acc := 0
	for _, seq := range seqs {
		acc += seq[(len(seq)-1)/2]
	}
	return acc
}

func read(fn string) ([][]int, [][]int, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	edges := [][]int{}
	seqs := [][]int{}
	s := bufio.NewScanner(f)
	addEdges := true
	for s.Scan() {
		line := string(s.Text())
		addEdges = addEdges && line != ""
		if addEdges {
			e := []int{}
			for _, s := range strings.Split(line, "|") {
				n, _ := strconv.Atoi(s)
				e = append(e, n)
			}
			edges = append(edges, e)
		} else {
			seq := []int{}
			for _, s := range strings.Split(line, ",") {
				n, _ := strconv.Atoi(s)
				seq = append(seq, n)
			}
			seqs = append(seqs, seq)
		}
	}
	return edges, seqs, nil
}

func main() {
	edges, seq, err := read(fn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot read file %s: %v\n", fn, err)
		os.Exit(1)
	}
	fmt.Printf("part 1: %v\n", one(edges, seq))
	fmt.Printf("part 2: %v\n", two(edges, seq))
}
