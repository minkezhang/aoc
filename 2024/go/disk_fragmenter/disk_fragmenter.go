package disk_fragmenter

import (
	"strconv"
)

type P struct{}

func (p P) Name() string  { return "2024/09" }
func (p P) Input() string { return "disk_fragmenter.txt" }

func (p P) F(data []byte) (int, int, error) {
	return one(data), two(data), nil
}

type file struct {
	id        int
	blocks    [][]int // [start, length]
	defragged [][]int
}

func genfiles(data []byte) (empty *file, files []*file) {
	empty = &file{id: -1, blocks: [][]int{}}
	files = []*file{}

	p := 0
	id := 0
	for i, d := range data {
		l, _ := strconv.Atoi(string(d))
		block := []int{p, l}
		if i%2 == 0 { // data
			files = append(files, &file{id: id, blocks: [][]int{block}})
			id += 1
		} else {
			empty.blocks = append(empty.blocks, block)
		}
		p += l
	}
	return empty, files
}

func two(data []byte) int {
	empty, files := genfiles(data)

	for i := len(files) - 1; i >= 0; i-- {
		f := files[i]
		for i, c := range empty.blocks {
			if len(f.blocks) == 0 {
				break
			}
			b := f.blocks[0] // An input file is always contiguous.
			if c[0] < b[0] && c[1] >= b[1] {
				f.defragged = append(f.defragged, []int{c[0], b[1]})
				c[1] -= b[1]
				c[0] += b[1]
				b[1] = 0
				f.blocks = f.blocks[:len(f.blocks)-1]
				if c[1] == 0 {
					empty.blocks = append(empty.blocks[:i], empty.blocks[i+1:]...)
				}
			}
		}
	}
	return checksum(files)
}

func one(data []byte) int {
	empty, files := genfiles(data)

	for i := len(files) - 1; i >= 0; i-- {
		f := files[i]
		for len(empty.blocks) > 0 && len(f.blocks) > 0 {
			b := f.blocks[len(f.blocks)-1]
			c := empty.blocks[0]
			if c[0] > b[0] {
				break
			}
			if b[1] <= c[1] {
				f.defragged = append(f.defragged, []int{c[0], b[1]})
				c[1] -= b[1]
				c[0] += b[1]
				b[1] = 0
			} else {
				f.defragged = append(f.defragged, []int{c[0], c[1]})
				b[1] -= c[1]
				c[1] = 0
			}
			if b[1] == 0 {
				f.blocks = f.blocks[:len(f.blocks)-1]
			}
			if c[1] == 0 {
				empty.blocks = empty.blocks[1:]
			}
		}
	}
	return checksum(files)
}

func checksum(fs []*file) int {
	sum := 0
	for _, f := range fs {
		blocks := append([][]int{}, f.blocks...)
		blocks = append(blocks, f.defragged...)
		for _, b := range blocks {
			for i := 0; i < b[1]; i++ {
				sum += f.id * (b[0] + i)
			}
		}
	}
	return sum
}
