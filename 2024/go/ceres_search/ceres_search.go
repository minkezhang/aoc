package ceres_search

import (
	"bufio"
	"bytes"
)

type P struct{}

func (p P) Name() string  { return "2024/04" }
func (p P) Input() string { return "ceres_search.txt" }

func (p P) F(data []byte) (int, int, error) {
	lines := []string{}
	buf := bytes.NewReader(data)
	s := bufio.NewScanner(buf)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	return one(lines[0 : len(lines)-1]), two(lines[0 : len(lines)-1]), nil
}

func one(data []string) int {
	l := len(data[0])
	acc := 0
	for i := 0; i < l-3; i++ {
		for j := 0; j < len(data); j++ {
			if data[j][i:i+4] == "XMAS" {
				acc += 1
			}
			if data[j][i:i+4] == "SAMX" {
				acc += 1
			}
		}
	}
	for i := 0; i < l; i++ {
		for j := 0; j < len(data)-3; j++ {
			if data[j][i] == 'X' && data[j+1][i] == 'M' && data[j+2][i] == 'A' && data[j+3][i] == 'S' {
				acc += 1
			}
			if data[j][i] == 'S' && data[j+1][i] == 'A' && data[j+2][i] == 'M' && data[j+3][i] == 'X' {
				acc += 1
			}
		}
	}
	for i := 0; i < l-3; i++ {
		for j := 0; j < len(data)-3; j++ {
			if data[j][i] == 'X' && data[j+1][i+1] == 'M' && data[j+2][i+2] == 'A' && data[j+3][i+3] == 'S' {
				acc += 1
			}
			if data[j][i] == 'S' && data[j+1][i+1] == 'A' && data[j+2][i+2] == 'M' && data[j+3][i+3] == 'X' {
				acc += 1
			}
			if data[j][i+3] == 'X' && data[j+1][i+2] == 'M' && data[j+2][i+1] == 'A' && data[j+3][i] == 'S' {
				acc += 1
			}
			if data[j][i+3] == 'S' && data[j+1][i+2] == 'A' && data[j+2][i+1] == 'M' && data[j+3][i] == 'X' {
				acc += 1
			}
		}
	}
	return acc
}

func two(data []string) int {
	l := len(data[0])
	acc := 0
	for i := 0; i < l-2; i++ {
		for j := 0; j < len(data)-2; j++ {
			if data[j+1][i+1] == 'A' && ((data[j][i] == 'M' && data[j+2][i+2] == 'S') || (data[j][i] == 'S' && data[j+2][i+2] == 'M')) && ((data[j][i+2] == 'M' && data[j+2][i] == 'S') || (data[j][i+2] == 'S' && data[j+2][i] == 'M')) {
				acc += 1
			}
		}
	}
	return acc
}
