package solutions

func twoSum(nums []int, target int) []int {
	targets := make([]int, 2)

	for index, value := range nums {
		rest := target - value

		for index2, value := range nums {
			if index != index2 && value == rest {
				targets[0] = index
				targets[1] = index2
				return targets
			}
		}
	}
}
