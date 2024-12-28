package historian_hysteria

import (
	"bytes"
	"encoding/csv"
	"math"
	"sort"
	"strconv"
)

type P struct{}

func (p P) Name() string  { return "2024/01" }
func (p P) Input() string { return "historian_hysteria.tsv" }

func (p P) F(data []byte) (int, int, error) {
	buf := bytes.NewReader(data)
	r := csv.NewReader(buf)
	r.Comma = ' '
	rs, err := r.ReadAll()
	if err != nil {
		return 0, 0, err
	}

	var as []int
	var bs []int
	for _, record := range rs {
		a, _ := strconv.Atoi(record[0])
		b, _ := strconv.Atoi(record[len(record)-1])
		as = append(as, a)
		bs = append(bs, b)
	}

	return one(as, bs), two(as, bs), nil
}

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
