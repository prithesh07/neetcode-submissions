import "slices"

type Result struct {
	a int
	b int
	c int
}

func threeSum(nums []int) [][]int {
	slices.Sort(nums)

	n := len(nums)

	result_set := make(map[Result]struct{})
	result := [][]int {}

	for i := 0; i < n; i++ {
		target := -1 * nums[i]

		left := i + 1
		right := n - 1

		for left < right {
			cur_sum := nums[left] + nums[right]

			if cur_sum == target {
				result_set [Result {nums[i], nums[left], nums[right]}] = struct{}{}
				left++
				right--
			} else if cur_sum < target {
				left++
			} else {
				right--
			}
		}

	}

	for key, _ := range result_set {
		result = append(result, []int {key.a, key.b, key.c})
	}

	return result
}
