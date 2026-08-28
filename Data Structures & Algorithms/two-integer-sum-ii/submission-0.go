func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		cur_sum := numbers[left] + numbers[right]

		if cur_sum == target {
			return []int {left+1, right + 1}
		} else if cur_sum < target {
			left++
		} else {
			right--
		}
	}

	return nil
}
