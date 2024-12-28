package mull_it_over

import (
	"regexp"
	"strconv"
	"strings"
)

type P struct{}

func (p P) Name() string  { return "2024/03" }
func (p P) Input() string { return "mull_it_over.txt" }

func (p P) F(data []byte) (int, int, error) { return one(string(data)), two(string(data)), nil }

func match(data string) [][]int {
	re := regexp.MustCompile(`(?m)mul\((?P<a>[0-9]{1,3}),(?P<b>[0-9]{1,3})\)`)
	matches := re.FindAllSubmatch([]byte(data), -1)
	pairs := [][]int{}
	for _, m := range matches {
		a, _ := strconv.Atoi(string(m[re.SubexpIndex("a")]))
		b, _ := strconv.Atoi(string(m[re.SubexpIndex("b")]))
		pairs = append(pairs, []int{a, b})
	}
	return pairs
}

func one(data string) int {
	pairs := match(data)
	acc := 0
	for _, p := range pairs {
		acc += p[0] * p[1]
	}
	return acc
}

func two(data string) int {
	donts := strings.Split(data, "don't()") // Every subsequent part is preceeded by a "don't()"
	acc := one(donts[0])
	if len(donts) > 1 {
		for _, dont := range donts[1:] {
			dos := strings.Split(dont, "do()") // Every subsequent part is preceeded by a "do()"
			if len(dos) > 1 {
				for _, do := range dos[1:] {
					acc += one(do)
				}
			}
		}
	}
	return acc
}
