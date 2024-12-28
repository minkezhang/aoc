package base

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type P[U, V any] interface {
	Name() string
	Input() string
	F(data []byte) (U, V, error)
}

func ReadOrDie(fn string) []byte {
	fp, err := os.Open(fn)
	if err != nil {
		panic(fmt.Sprintf("cannot open file \"%s\": %v", fn, err))
	}
	defer fp.Close()

	data, err := io.ReadAll(fp)
	if err != nil {
		panic(fmt.Sprintf("cannot read file \"%s\": %v", fn, err))
	}
	return data

}

type T[U, V any] struct {
	P P[U, V]

	data        []byte
	initialized bool
}

func (t *T[U, V]) ReadOrDie() {
	if t.initialized {
		return
	}
	t.initialized = true
	t.data = ReadOrDie(filepath.Join("..", "inputs", t.P.Input()))
}

func (t *T[U, V]) Solve() (U, V, error) {
	t.ReadOrDie()
	return t.P.F(t.data)
}

func (t *T[U, V]) Print() {
	u, v, err := t.Solve()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: cannot solve problem: %v", t.P.Name(), err)
	}

	fmt.Printf("%s part 1: %v\n", t.P.Name(), u)
	fmt.Printf("%s part 2: %v\n", t.P.Name(), v)
}
