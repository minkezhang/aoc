package main

import (
	"regexp"
	"strconv"
	"fmt"
	"io"
	"os"
)

const (
	fn = "mull.txt"
)

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

func read(fn string) (string, error) {
	f, err := os.Open(fn)
	if err != nil {
		return "", err
	}
	defer f.Close()

	s, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(s), nil
}

func main() {
	data, err := read(fn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot read file %s: %v\n", fn, err)
		os.Exit(1)
	}

	fmt.Printf("part 1: %v\n", one(data))
}
