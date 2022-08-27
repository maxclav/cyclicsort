package cyclicsort

// Sort sorts the slice `nums` if it is valid or return an error if not.
// It is valid if each items are uniques and in the range of [0, N-1]
// where `N` is the len(nums).
func Sort(nums []int) error {
	if err := Validate(nums); err != nil {
		return err
	}
	if len(nums) <= 1 {
		return nil
	}
	SortValid(nums)
	return nil
}

// MustSort sort or panic the slice.
func MustSort(nums []int) {
	if err := Sort(nums); err != nil {
		panic(err)
	}
}

// SortValid sorts a valid slice `nums`.
// It is valid if each items are uniques and in the range of [0, N-1]
// where `N` is the len(nums).
// Do NOT use if the slice if not valid.
func SortValid(nums []int) {
	if len(nums) <= 1 {
		return
	}
	var index int = 0
	for index < len(nums)-1 {
		if index != nums[index] {
			nums[index], nums[nums[index]] = nums[nums[index]], nums[index]
		} else {
			index++
		}
	}
}

// Validate returns an error if the slice is not valid for a cyclic sort.
// It is valid if each items are uniques and in the range of [0, N-1]
// where `N` is the len(nums).
func Validate(nums []int) error {
	m := make(map[int]struct{})
	for _, num := range nums {
		if num < 0 {
			return ErrValueNegative
		}
		if num > len(nums)-1 {
			return ErrValueBig
		}
		if _, ok := m[num]; ok {
			return ErrValueDuplicated
		}
		m[num] = struct{}{}
	}
	return nil
}
