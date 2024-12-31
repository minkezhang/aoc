package disk_fragmenter

import (
	"testing"
)

func TestChecksum(t *testing.T) {
	for _, c := range []struct {
		name string
		fs   []*file
		want int
	}{
		{
			name: "Trivial",
			fs:   nil,
			want: 0,
		},
		{
			name: "Simple",
			fs: []*file{
				&file{
					id:     1,
					blocks: [][]int{[]int{1, 1}},
				},
			},
			want: 1,
		},
		{
			name: "Simple/WithDefragged",
			fs: []*file{
				&file{
					id:        1,
					blocks:    [][]int{[]int{1, 1}},
					defragged: [][]int{[]int{2, 1}},
				},
			},
			want: 3,
		},
	} {
		t.Run(c.name, func(t *testing.T) {
			if got := checksum(c.fs); got != c.want {
				t.Errorf("checksum() = %v, want = %v", got, c.want)
			}
		})
	}
}

func TestTwo(t *testing.T) {
	for _, c := range []struct {
		name string
		data []byte
		want int
	}{
		{
			name: "Example",
			data: []byte("2333133121414131402"),
			want: 2858,
		},
	} {
		t.Run(c.name, func(t *testing.T) {
			if got := two(c.data); got != c.want {
				t.Errorf("two() = %v, want = %v", got, c.want)
			}
		})
	}
}

func TestOne(t *testing.T) {
	for _, c := range []struct {
		name string
		data []byte
		want int
	}{
		{
			name: "Trivial",
			data: []byte(""),
			want: 0,
		},
		{
			name: "Simple/233",
			data: []byte("233"),
			want: 9,
		},
		{
			name: "Example/12345",
			data: []byte("12345"),
			want: 2*1 + 2*2 + 1*3 + 1*4 + 1*5 + 2*6 + 2*7 + 2*8,
		},
		{
			name: "Example/2333133121414131402",
			data: []byte("2333133121414131402"),
			want: 1928,
		},
	} {
		t.Run(c.name, func(t *testing.T) {
			if got := one(c.data); got != c.want {
				t.Errorf("one() = %v, want = %v", got, c.want)
			}
		})
	}
}
