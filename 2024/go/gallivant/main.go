package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	fn = "gallivant.txt"
)

// two detects how many cycles exist by adding a single obstacle.
//
// TODO(minkezhang): Only add obstacle along the original path instead of
// complete brute force.
func two(data string) int {
	acc := 0
	for i := 0; i < len(data); i++ {
		if data[i] != '#' && data[i] != '^' && data[i] != '\n' {
			test := []rune(data)
			test[i] = '#'
			if isLoop := one(string(test)) == -1; isLoop {
				acc += 1
			}
		}
	}
	return acc
}

func one(data string) int {
	m := [][]bool{}         // m[X][Y] == true --> blocked
	v := [][]map[int]bool{} // Track visited nodes.
	var start []int
	for y, line := range strings.Split(data, "\n") {
		if len(line) == 0 {
			continue
		}
		r := make([]bool, len(line))
		for x, c := range line {
			if c == '#' {
				r[x] = true
			}
			if c == '^' {
				start = []int{x, y}
			}

		}
		m = append(m, r)
		vr := make([]map[int]bool, len(line))
		for i := range vr {
			vr[i] = map[int]bool{}
		}
		v = append(v, vr)
	}

	acc := 1
	direction := 0 // N E S W
	v[start[1]][start[0]][direction] = true
	l := len(m[0])
	pos := start
	for {
		x, y := pos[0], pos[1]
		var next []int
		switch direction {
		case 0:
			next = []int{x, y - 1}
		case 1:
			next = []int{x + 1, y}
		case 2:
			next = []int{x, y + 1}
		case 3:
			next = []int{x - 1, y}
		default:
			panic("invalid direction")
		}
		x, y = next[0], next[1]
		if x < 0 || x >= l || y < 0 || y >= len(m) { // Out of bounds.
			break
		}
		if v[y][x][direction] { // Detected cycle.
			return -1
		}
		if m[y][x] {
			direction = (direction + 1) % 4
		} else {
			pos = next
			if len(v[y][x]) == 0 { // Avoid double counting
				acc++
				v[y][x][direction] = true
			}
		}
	}
	return acc
}

func read(fn string) (string, error) {
	f, err := os.Open(fn)
	if err != nil {
		return "", err
	}
	defer f.Close()
	data, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func main() {
	data, err := read(fn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot read file %s: %v\n", fn, err)
		os.Exit(1)
	}
	fmt.Printf("part 1: %v\n", one(data))
	fmt.Printf("part 2: %v\n", two(data))
}
