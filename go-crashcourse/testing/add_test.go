package add_test

import (
	"testing"

	"github.com/nii236/gotraining/go-crashcourse/testing"
)

var tests = []struct {
	in  []int
	out int
}{
	{[]int{1, 2}, 3},
	{[]int{-1, 2}, 1},
	{[]int{1, 0}, 1},
}

func TestFunc(t *testing.T) {
	for _, tt := range tests {
		got := add.Do(tt.in[0], tt.in[1])
		expected := tt.out
		if got != tt.out {
			t.Errorf("Error! Got %d, expected %d", got, expected)
		}

	}
}
