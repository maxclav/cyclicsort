package cyclicsort

// SortValid sorts a valid slice `nums`.
// It is valid if each items are uniques and in the range of [0, N-1]
// where `N` is the len(nums).
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

// Sort sorts the slice `nums` if it is valid.
// It is valid if each items are uniques and in the range of [0, N-1]
// where `N` is the len(nums).
func Sort(nums []int) {
	if len(nums) <= 1 || !IsValid(nums) {
		return
	}
	SortValid(nums)
}

// IsValid returns whether the slice `nums`
// is valid for the cyclic sort (true) or not (false).
//
// It is valid if each items are uniques and in the range of [0, N-1]
// where `N` is the len(nums).
func IsValid(nums []int) bool {
	m := make(map[int]struct{})
	for _, num := range nums {
		if num > len(nums) {
			return false
		}
		if _, ok := m[num]; ok {
			return false
		}
		m[num] = struct{}{}
	}
	return true
}
