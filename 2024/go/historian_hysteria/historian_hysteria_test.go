package historian_hysteria

import (
	"testing"
)

func TestTwo(t *testing.T) {
	as := []int{3, 4, 2, 1, 3, 3}
	bs := []int{4, 3, 5, 3, 9, 3}
	want := 31
	if got := two(as, bs); got != want {
		t.Errorf("two() = %v, want = %v", got, want)
	}
}
