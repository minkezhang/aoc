package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	fn = "rednose.tsv"
)

func read(fn string) ([][]int, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reports := [][]int{}
	s := bufio.NewScanner(f)

	for s.Scan() {
		if line := string(s.Text()); line != "" {
			levels := []int{}
			for _, l := range strings.Split(line, " ") {
				level, _ := strconv.Atoi(l)
				levels = append(levels, level)
			}
			reports = append(reports, levels)
		}
	}
	return reports, nil
}

func safe(report []int) bool {
	if len(report) == 1 {
		return true
	}

	increase := report[0] < report[1]

	for i := range report[:len(report)-1] {
		if increase != (report[i] < report[i+1]) {
			return false
		}
		diff := report[i] - report[i+1]
		if diff < 0 {
			diff *= -1
		}
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func one(reports [][]int) int {
	acc := 0
	for _, r := range reports {
		if safe(r) {
			acc += 1
		}
	}
	return acc
}

func damp(report []int) bool {
	if len(report) == 1 {
		return true
	}

	for i := 0; i < len(report); i++ {
		candidate := []int{}
		if i > 0 {
			candidate = append(candidate, report[0:i]...)
		}
		if i < len(report)-1 {
			candidate = append(candidate, report[i+1:len(report)]...)
		}
		if safe(candidate) {
			return true
		}
	}

	return false
}

func two(reports [][]int) int {
	acc := 0
	for _, r := range reports {
		if damp(r) {
			acc += 1
		}
	}
	return acc
}

func main() {
	reports, err := read(fn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot read file %s: %v\n", fn, err)
		os.Exit(1)
	}
	fmt.Printf("part 1: %v\n", one(reports))
	fmt.Printf("part 2: %v\n", two(reports))
}
