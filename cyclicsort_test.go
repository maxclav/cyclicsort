package cyclicsort_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "github.com/maxclav/cyclicsort"
)

// TestUnit_IsValid is a unit test
// for the IsValid() function.
func TestUnit_IsValid(t *testing.T) {
	type test struct {
		input []int
		want  bool
	}

	tests := map[string]test{
		"valid nil slice": {
			[]int{}, true,
		},
		"valid one item slice": {
			[]int{1}, true,
		},
		"invalid one item slice": {
			[]int{2}, false,
		},
		"valid slice": {
			[]int{1, 2, 3, 4, 5}, true,
		},
		"invalid slice (because of first item)": {
			[]int{10, 2, 3, 4, 5}, false,
		},
		"invalid slice (because of last item)": {
			[]int{1, 2, 3, 4, 10}, false,
		},
		"invalid slice (because of all items)": {
			[]int{10, 20, 30, 40, 50}, false,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			var got, want bool
			defer func() {
				p := recover()
				require.Nil(t, p)
				assert.Equal(t, want, got)
			}()
			want = test.want
			got = IsValid(test.input)
		})
	}
}

// TestUnit_Sort is a unit tests
// for the SortValid() function.
func TestUnit_SortValid(t *testing.T) {
	type test struct {
		input, want []int
	}

	tests := map[string]test{
		"nil slice": {
			[]int{}, []int{},
		},
		"one item slice": {
			[]int{0}, []int{0},
		},
		"already sorted": {
			[]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 3, 4},
		},
		"random sorted": {
			[]int{3, 1, 0, 2, 4}, []int{0, 1, 2, 3, 4},
		},
		"descending sorted": {
			[]int{4, 3, 2, 1, 0}, []int{0, 1, 2, 3, 4},
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			var got, want []int
			defer func() {
				p := recover()
				require.Nil(t, p)
				assert.Equal(t, want, got)
			}()
			want = test.want
			Sort(test.input)
			got = test.input
		})
	}
}
