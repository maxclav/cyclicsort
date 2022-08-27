package cyclicsort_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "github.com/maxclav/cyclicsort"
)

// TestUnit_Sort is a unit tests
// for the Sort() function.
func TestUnit_Sort(t *testing.T) {
	type want struct {
		nums []int
		err  bool
	}
	type test struct {
		input []int
		want  want
	}

	tests := map[string]test{
		"valid slice nil slice": {
			[]int{}, want{
				nums: []int{},
				err:  false,
			},
		},
		"valid slice one item slice": {
			[]int{0}, want{
				nums: []int{0},
				err:  false,
			},
		},
		"valid two items slice sorted": {
			[]int{0, 1}, want{
				nums: []int{0, 1},
				err:  false,
			},
		},
		"valid two items slice unsorted": {
			[]int{1, 0}, want{
				nums: []int{0, 1},
				err:  false,
			},
		},
		"valid slice already sorted": {
			[]int{0, 1, 2, 3, 4}, want{
				nums: []int{0, 1, 2, 3, 4},
				err:  false,
			},
		},
		"valid slice random sorted": {
			[]int{3, 1, 0, 2, 4}, want{
				nums: []int{0, 1, 2, 3, 4},
				err:  false,
			},
		},
		"valid slice descending sorted": {
			[]int{4, 3, 2, 1, 0}, want{
				nums: []int{0, 1, 2, 3, 4},
				err:  false,
			},
		},
		"invalid first item (< 0)": {
			[]int{-3, 1, 0, 2, 4}, want{
				nums: nil,
				err:  true,
			},
		},
		"invalid first item (<<< 0)": {
			[]int{math.MinInt, 1, 0, 2, 4}, want{
				nums: nil,
				err:  true,
			},
		},
		"invalid first item (> max)": {
			[]int{10, 1, 0, 2, 4}, want{
				nums: nil,
				err:  true,
			},
		},
		"invalid first item (>>> max)": {
			[]int{math.MaxInt, 1, 0, 2, 4}, want{
				nums: nil,
				err:  true,
			},
		},
		"invalid slice same values": {
			[]int{1, 1, 1, 1, 1}, want{
				nums: nil,
				err:  true,
			},
		},
		"invalid slice same values (zeros)": {
			[]int{0, 0, 0, 0, 0}, want{
				nums: nil,
				err:  true,
			},
		},
		"invalid slice same values (max)": {
			[]int{4, 4, 4, 4, 4}, want{
				nums: nil,
				err:  true,
			},
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			var got, want []int
			var gotErr, wantErr bool
			defer func() {
				p := recover()
				require.Nil(t, p)
				assert.Equal(t, wantErr, gotErr)
				if !wantErr {
					assert.Equal(t, want, got)
				}
			}()
			want = test.want.nums
			wantErr = test.want.err
			got = test.input
			err := Sort(got)
			gotErr = err != nil
		})
	}
}

// TestUnit_MustSort is a unit tests
// for the MustSort() function.
func TestUnit_MustSort(t *testing.T) {
	type want struct {
		nums []int
		err  bool
	}
	type test struct {
		input []int
		want  want
	}

	tests := map[string]test{
		"valid slice nil slice": {
			[]int{}, want{
				nums: []int{},
				err:  false,
			},
		},
		"valid slice one item slice": {
			[]int{0}, want{
				nums: []int{0},
				err:  false,
			},
		},
		"valid two items slice sorted": {
			[]int{0, 1}, want{
				nums: []int{0, 1},
				err:  false,
			},
		},
		"valid two items slice unsorted": {
			[]int{1, 0}, want{
				nums: []int{0, 1},
				err:  false,
			},
		},
		"valid slice already sorted": {
			[]int{0, 1, 2, 3, 4}, want{
				nums: []int{0, 1, 2, 3, 4},
				err:  false,
			},
		},
		"valid slice random sorted": {
			[]int{3, 1, 0, 2, 4}, want{
				nums: []int{0, 1, 2, 3, 4},
				err:  false,
			},
		},
		"valid slice descending sorted": {
			[]int{4, 3, 2, 1, 0}, want{
				nums: []int{0, 1, 2, 3, 4},
				err:  false,
			},
		},
		"invalid first item (< 0)": {
			[]int{-3, 1, 0, 2, 4}, want{
				nums: nil,
				err:  true,
			},
		},
		"invalid first item (<<< 0)": {
			[]int{math.MinInt, 1, 0, 2, 4}, want{
				nums: nil,
				err:  true,
			},
		},
		"invalid first item (> max)": {
			[]int{10, 1, 0, 2, 4}, want{
				nums: nil,
				err:  true,
			},
		},
		"invalid first item (>>> max)": {
			[]int{math.MaxInt, 1, 0, 2, 4}, want{
				nums: nil,
				err:  true,
			},
		},
		"invalid slice same values": {
			[]int{1, 1, 1, 1, 1}, want{
				nums: nil,
				err:  true,
			},
		},
		"invalid slice same values (zeros)": {
			[]int{0, 0, 0, 0, 0}, want{
				nums: nil,
				err:  true,
			},
		},
		"invalid slice same values (max)": {
			[]int{4, 4, 4, 4, 4}, want{
				nums: nil,
				err:  true,
			},
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			var got, want []int
			var wantErr bool
			defer func() {
				p := recover()
				if wantErr {
					assert.NotNil(t, p)
				} else {
					require.Nil(t, p)
					assert.Equal(t, want, got)
				}
			}()
			want = test.want.nums
			wantErr = test.want.err
			got = test.input
			MustSort(got)
		})
	}
}

// TestUnit_SortValid is a unit tests
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
		"two items slice sorted": {
			[]int{0, 1}, []int{0, 1},
		},
		"two items slice unsorted": {
			[]int{1, 0}, []int{0, 1},
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
			SortValid(test.input)
			got = test.input
		})
	}
}

// TestUnit_IsValid is a unit test
// for the IsValid() function.
func TestUnit_IsValid(t *testing.T) {
	type test struct {
		input   []int
		wantErr bool
	}

	tests := map[string]test{
		"valid nil slice": {
			[]int{}, false,
		},
		"valid one item slice (zero value)": {
			[]int{0}, false,
		},
		"invalid one item slice (<<< 0)": {
			[]int{math.MinInt}, true,
		},
		"invalid one item slice (negative)": {
			[]int{-1}, true,
		},
		"invalid one item slice (> max)": {
			[]int{1}, true,
		},
		"invalid one item slice (>>> max)": {
			[]int{math.MaxInt}, true,
		},
		"valid slice (all different values ascending)": {
			[]int{0, 1, 2, 3, 4}, false,
		},
		"valid slice (all different values descending)": {
			[]int{4, 3, 2, 1, 0}, false,
		},
		"invalid slice (all same values)": {
			[]int{2, 2, 2, 2, 2}, true,
		},
		"invalid slice (all zero values)": {
			[]int{0, 0, 0, 0, 0}, true,
		},
		"invalid slice (all max values)": {
			[]int{4, 4, 4, 4, 4}, true,
		},
		"invalid slice (because of first item < 0)": {
			[]int{-1, 2, 3, 4, 5}, true,
		},
		"invalid slice (because of first item > max)": {
			[]int{10, 2, 3, 4, 5}, true,
		},
		"invalid slice (because of last item < 0)": {
			[]int{1, 2, 3, 4, -1}, true,
		},
		"invalid slice (because of last item > max)": {
			[]int{1, 2, 3, 4, 10}, true,
		},
		"invalid slice (because of all items < 0)": {
			[]int{-10, -20, -30, -40, -50}, true,
		},
		"invalid slice (because of all items > max)": {
			[]int{10, 20, 30, 40, 50}, true,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			var gotErr, wantErr bool
			defer func() {
				p := recover()
				require.Nil(t, p)
				assert.Equal(t, wantErr, gotErr)
			}()
			wantErr = test.wantErr
			err := Validate(test.input)
			gotErr = err != nil
		})
	}
}
