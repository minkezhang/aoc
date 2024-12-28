package red_nosed_reports

import (
	"testing"
)

func TestDamp(t *testing.T) {
	configs := []struct {
		name   string
		report []int
		want   bool
	}{
		{
			name:   "Remove/i=1",
			report: []int{1, 3, 2, 4, 5},
			want:   true,
		},
	}

	for _, c := range configs {
		t.Run(c.name, func(t *testing.T) {
			if got := damp(c.report); got != c.want {
				t.Errorf("damp() = %v, want = %v", got, c.want)
			}
		})
	}
}
