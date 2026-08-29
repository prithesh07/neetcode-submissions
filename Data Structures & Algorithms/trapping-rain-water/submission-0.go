func trap(height []int) int {
	n := len(height)

	left_max := make([]int, n)
	right_max := make([]int, n)

	left_max[0] = height[0]
	for i:=1; i<n; i++ {
		left_max[i] = max(left_max[i-1], height[i])
	}
 
	right_max[n-1] = height[n-1]
	for i:=n-2; i>=0; i-- {
		right_max[i] = max(right_max[i+1], height[i])
	}

	total_water := 0
	for i:=1; i<n-1; i++ {
		total_water += max(0, min(left_max[i-1], right_max[i+1]) - height[i])
	}	

	return total_water
}
