func lengthOfLongestSubstring(s string) int {
	left := 0
	right := 0
	hash_set := make(map[rune]int)
	result := 0
	n := len(s)

	for right < n {
		if val, exists := hash_set[rune(s[right])]; exists {
			for left <= val {
				delete(hash_set, rune(s[left]))
				left++
			}
		} else {
			hash_set[rune(s[right])] = right
			result = max(result, right - left + 1)
			right++
		}
	}

	return result
}
