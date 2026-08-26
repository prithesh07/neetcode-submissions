func topKFrequent(nums []int, k int) []int {
	hash_map := make(map[int]int)

	for _, val := range nums {
		hash_map[val] += 1
	}

	list := make([][]int, len(nums)+1)
	for key, value := range hash_map {
		list[value] = append(list[value], key)
	}

	result := []int{}

	for i:=len(nums); (i >=0 && k > 0); i-- {
		for _, val := range(list[i]) {
			k--
			result = append(result, val)

			if k <= 0 {
				break;
			}
		}
	}

	return result
}
