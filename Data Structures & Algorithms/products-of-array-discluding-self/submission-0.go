func productExceptSelf(nums []int) []int {
	length := len(nums)
	result := make([]int, length)

	result[0] = 1
	for i:=1; i<length; i++ {
		result[i] = nums[i-1] * result[i-1]
	}

	run := nums[length-1]

	for i:=length-2; i>=0; i-- {
		result[i] *= run
		run *= nums[i] 
	}

	return result
}
