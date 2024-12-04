package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

const (
	fn = "historian.tsv"
)

func one(as []int, bs []int) int {
	sort.Ints(as)
	sort.Ints(bs)
	diff := 0
	for i, a := range as {
		diff += int(math.Abs(float64(a - bs[i])))
	}
	return diff
}

func two(as []int, bs []int) int {
	acounter := map[int]int{}
	bcounter := map[int]int{}

	for _, a := range as {
		if _, ok := acounter[a]; !ok {
			acounter[a] = 1
		} else {
			acounter[a] += 1
		}
	}
	for _, b := range bs {
		if _, ok := bcounter[b]; !ok {
			bcounter[b] = 1
		} else {
			bcounter[b] += 1
		}
	}

	diff := 0
	for a, n := range acounter {
		if m, ok := bcounter[a]; ok {
			diff += a * n * m
		}
	}
	return diff
}

func read(fn string) ([]int, []int, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ' '
	rs, err := r.ReadAll()
	if err != nil {
		return nil, nil, err
	}

	var as []int
	var bs []int
	for _, record := range rs {
		a, _ := strconv.Atoi(record[0])
		b, _ := strconv.Atoi(record[len(record)-1])
		as = append(as, a)
		bs = append(bs, b)
	}
	return as, bs, nil
}
func main() {
	as, bs, err := read(fn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot read file %s: %v\n", fn, err)
	}

	fmt.Printf("part 1: %v\n", one(as, bs))
	fmt.Printf("part 2: %v\n", two(as, bs))
}
