package resonant_collinearity

import (
	"strings"
)

type M struct {
	dim       [2]int            // {x, y}
	nodes     map[rune][][2]int // node -> [{x, y}]
	antinodes map[[2]int]bool   // [x, y] -> true
}

func New(data []byte) *M {
	lines := strings.Split(string(data), "\n")
	height := 0

	nodes := map[rune][][2]int{}
	for y, l := range lines {
		if l != "" {
			height += 1
			for x, n := range l {
				if n != '.' {
					nodes[n] = append(nodes[n], [2]int{x, y})
				}
			}
		}
	}

	return &M{
		dim:   [2]int{len(lines[0]), height},
		nodes: nodes,
	}
}

func (m *M) One() int {
	m.antinodes = map[[2]int]bool{}
	for _, nodes := range m.nodes {
		for i, s := range nodes {
			for _, t := range nodes[i+1:] {
				d := [2]int{s[0] - t[0], s[1] - t[1]}
				candidates := [][2]int{
					[2]int{s[0] + d[0], s[1] + d[1]}, // s +  d
					[2]int{t[0] - d[0], t[1] - d[1]}, // t + -d
				}
				for _, c := range candidates {
					if c[0] >= 0 && c[0] < m.dim[0] && c[1] >= 0 && c[1] < m.dim[1] {
						m.antinodes[c] = true
					}
				}
			}
		}
	}
	return len(m.antinodes)
}

func (m *M) Two() int {
	m.antinodes = map[[2]int]bool{}
	for _, nodes := range m.nodes {
		for i, s := range nodes {
			for _, t := range nodes[i+1:] {
				d := [2]int{s[0] - t[0], s[1] - t[1]}

				// s + d
				c := s // Original antenna is also considered an antinode.
				for c[0] >= 0 && c[0] < m.dim[0] && c[1] >= 0 && c[1] < m.dim[1] {
					m.antinodes[c] = true
					c = [2]int{c[0] + d[0], c[1] + d[1]}
				}
				// t + -d
				c = t
				for c[0] >= 0 && c[0] < m.dim[0] && c[1] >= 0 && c[1] < m.dim[1] {
					m.antinodes[c] = true
					c = [2]int{c[0] - d[0], c[1] - d[1]}
				}
			}
		}
	}
	return len(m.antinodes)
}

func (m *M) String() string {
	lines := [][]byte{}
	for y := 0; y < m.dim[1]; y++ {
		l := []byte{}
		for x := 0; x < m.dim[0]; x++ {
			l = append(l, byte('.'))
		}
		lines = append(lines, l)
	}
	for a := range m.antinodes {
		lines[a[1]][a[0]] = byte('#')
	}
	for r, nodes := range m.nodes {
		for _, n := range nodes {
			if lines[n[1]][n[0]] == '#' {
				lines[n[1]][n[0]] = byte('*')
			} else {
				lines[n[1]][n[0]] = byte(r)
			}
		}
	}
	elems := []string{}
	for _, l := range lines {
		elems = append(elems, string(l))
	}
	return strings.Join(elems, "\n")
}

type P struct {}

func (p P) Name() string { return "2024/08" }
func (p P) Input() string { return "resonant_collinearity.txt" }

func (p P) F(data []byte) (int, int, error) { return New(data).One(), New(data).Two(), nil }
