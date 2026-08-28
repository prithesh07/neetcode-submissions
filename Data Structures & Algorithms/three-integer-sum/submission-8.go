import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)

	n := len(nums)
	result := [][]int{}

	for i := 0; i < n; i++ {
		// Since array is sorted, no possible triplet after this
		if nums[i] > 0 {
			break
		}

		// Skip duplicate first numbers
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := n - 1

		for left < right {
			target := nums[i] + nums[left] + nums[right]

			if target == 0 {
				result = append(result, []int{
					nums[i],
					nums[left],
					nums[right],
				})

				left++
				right--

				// Skip duplicate left values
				for left < right && nums[left] == nums[left-1] {
					left++
				}

				// Skip duplicate right values
				for left < right && nums[right] == nums[right+1] {
					right--
				}

			} else if target < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}
