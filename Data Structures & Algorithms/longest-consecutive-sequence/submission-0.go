func longestConsecutive(nums []int) int {
	all_values := make(map[int]struct{})
	visited := make(map[int]int)

	for _, val := range nums {
		all_values[val] = struct{}{}
	}

	max_length := 0
	for _, val := range nums {
		if _, exists := visited[val]; exists {
			continue
		}

		source := val
		visited[source] = 1
		length := 1

		for {
			if _, exist := all_values[val+1]; exist {
				if l, exists := visited[val+1]; exists {
					length += l
					break
				}

				val += 1
				visited[val] = 1
				length += 1

			} else {
				break
			}
		}

		visited[source] = length

		max_length = max(length, max_length)
	}

	return max_length
}
