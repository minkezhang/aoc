package red_nosed_reports

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"
)

type P struct{}

func (p P) Name() string  { return "2024/02" }
func (p P) Input() string { return "historian_hysteria.tsv" }

func (p P) F(data []byte) (int, int, error) {
	reports := [][]int{}
	buf := bytes.NewReader(data)
	s := bufio.NewScanner(buf)

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

	return one(reports), two(reports), nil
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
