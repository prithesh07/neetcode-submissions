func longestConsecutive(nums []int) int {
	all_values := make(map[int]struct{})

	for _, val := range nums {
		all_values[val] = struct{}{}
	}

	max_length := 0

	for key, _ := range all_values {

		if _, exists := all_values[key - 1]; !exists {

			length := 1

			for {
				if _, exist := all_values[key + length ]; exist {
					length++
				} else {
					break
				}
			}

			max_length = max(length, max_length)
		}
		
	}

	return max_length
}
