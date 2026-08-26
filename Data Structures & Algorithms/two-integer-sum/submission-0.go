func twoSum(nums []int, target int) []int {
	hash_map := make(map[int]int)

	for i, val := range nums {
		if val, exists := hash_map[target - val]; exists {
			return []int{val, i}
		}

		hash_map[val] = i
	}

	return nil
}
