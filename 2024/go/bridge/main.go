package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	fn = "bridge.txt"
)

// checkConcat if a target T may be made by two values X, Y s.t. T == X || Y
//
// Cases:
//
//	T: 12, Y: 2  --> Success
//	T: 12, Y: 1  --> Fail
//	T: 12, Y: 12 --> Success
//
// T: 12, Y: 12 reduces to T: 1, Y: 1 i.e. mod 10 and compare values.
func checkConcat(target int, x int) (int, bool) {
	if target == 0 && x != 0 {
		return 0, false
	}
	if target%10 == x%10 {
		if x < 10 {
			return target / 10, true
		} else {
			return checkConcat(target/10, x/10)
		}
	}
	return 0, false
}

func checkTwo(target int, xs []int) bool {
	if len(xs) == 0 {
		return false
	}
	if len(xs) == 1 {
		return target == xs[0]
	}
	x := xs[len(xs)-1]
	res, succ := checkConcat(target, x)
	return ((target >= x && checkTwo(target-x, xs[:len(xs)-1])) || (target%x == 0 && checkTwo(target/x, xs[:len(xs)-1]))) || (succ && checkTwo(res, xs[:len(xs)-1]))
}

func two(data []byte) int {
	acc := 0
	for _, l := range strings.Split(string(data), "\n") {
		if l == "" {
			continue
		}
		ps := strings.SplitN(l, ":", 2)
		qs := []string{ps[0]}
		qs = append(qs, strings.Split(strings.Trim(ps[1], " "), " ")...)
		xs := []int{}
		for _, p := range qs {
			x, _ := strconv.Atoi(p)
			xs = append(xs, x)
		}
		if checkTwo(xs[0], xs[1:]) {
			acc += xs[0]
		}
	}
	return acc
}

func check(target int, xs []int) bool {
	if len(xs) == 0 {
		return false
	}

	if len(xs) == 1 {
		return target == xs[0]
	}

	x := xs[len(xs)-1]
	return ((target >= x && check(target-x, xs[:len(xs)-1])) || (target%x == 0 && check(target/x, xs[:len(xs)-1])))
}

func one(data []byte) int {
	acc := 0
	for _, l := range strings.Split(string(data), "\n") {
		if l == "" {
			continue
		}
		ps := strings.SplitN(l, ":", 2)
		qs := []string{ps[0]}
		qs = append(qs, strings.Split(strings.Trim(ps[1], " "), " ")...)
		xs := []int{}
		for _, p := range qs {
			x, _ := strconv.Atoi(p)
			xs = append(xs, x)
		}
		if check(xs[0], xs[1:]) {
			acc += xs[0]
		}
	}
	return acc
}

func read(fn string) ([]byte, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return io.ReadAll(f)
}

func main() {
	data, err := read(fn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening file %s: %v\n", fn, err)
		os.Exit(1)
	}

	fmt.Printf("part 1: %v\n", one(data))
	fmt.Printf("part 2: %v\n", two(data))
}
