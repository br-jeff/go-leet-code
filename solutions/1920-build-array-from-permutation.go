package solutions

func buildArray(nums []int) []int {
	n := len(nums)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = nums[nums[i]]
	}

	return arr
}
