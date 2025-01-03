package guard_gallivant

import (
	"io"
	"os"
	"strings"
)

// two detects how many cycles exist by adding a single obstacle.
//
// TODO(minkezhang): Only add obstacle along the original path instead of
// complete brute force.
func two(data string) int {
	l := len(strings.Split(data, "\n")[0])
	p := path(data)
	acc := 0
	v := map[int]map[int]bool{}
	for _, pos := range p[1:] { // Skip initial guard position.
		if _, ok := v[pos[1]]; !ok {
			v[pos[1]] = map[int]bool{}
		}
		if v[pos[1]][pos[0]] {
			continue
		}
		v[pos[1]][pos[0]] = true
		test := []rune(data)
		test[pos[1]*(l+1)+pos[0]] = '#'
		if isLoop := one(string(test)) == -1; isLoop {
			acc += 1
		}
	}
	return acc
}

func path(data string) [][]int {
	p := [][]int{}
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

	direction := 0 // N E S W
	v[start[1]][start[0]][direction] = true
	l := len(m[0])
	pos := start
	for {
		p = append(p, pos)

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
			return nil
		}
		if m[y][x] {
			direction = (direction + 1) % 4
		} else {
			pos = next
			if len(v[y][x]) == 0 { // Avoid double counting
				v[y][x][direction] = true
			}
		}
	}
	return p
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

type P struct{}

func (p P) Name() string  { return "2024/06" }
func (p P) Input() string { return "guard_gallivant.txt" }

func (p P) F(data []byte) (int, int, error) { return one(string(data)), two(string(data)), nil }

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
