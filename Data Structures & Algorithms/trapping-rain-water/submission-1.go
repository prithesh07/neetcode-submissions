func trap(height []int) int {
	n := len(height)
	left := 1
	right := n - 2
	left_max := height[0]
	right_max := height[n-1]

	total_water := 0

	for left <= right {
		if left_max < right_max {
			total_water += max(0, (left_max - height[left]))
			left_max = max(left_max, height[left])
			left++
		} else {
			total_water += max(0, (right_max - height[right]))
			right_max = max(right_max, height[right])
			right--
		}
	}

	return total_water
}
