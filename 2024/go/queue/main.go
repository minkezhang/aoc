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

func one(edges [][]int, seqs [][]int) int {
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
	acc := 0
	for _, seq := range seqs {
		if v(seq) {
			acc += seq[(len(seq)-1)/2]
		}
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
}
