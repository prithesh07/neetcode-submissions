func checkValid(char_arr []int) int {
	total := 0
	max_val := math.MinInt

	for i:= 0; i<26; i++ {
		total += char_arr[i]
		max_val = max(max_val, char_arr[i])
	}

	return total - max_val
}

func characterReplacement(s string, k int) int {
	left := 0
	right := 0
	n := len(s)
	result := 0

	char_arr := make([]int, 26)

	for right < n {
		char_arr[s[right]-'A'] += 1

		if checkValid(char_arr) <= k {
			result = max(result, right - left + 1)
		} else {
			for checkValid(char_arr) > k {
				char_arr[s[left]-'A'] -= 1
				left++
			}
		}

		right++
	}

	return result
}
