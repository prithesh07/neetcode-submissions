func isAnagram(s string, t string) bool {
	hash_map := make(map[rune]int)

	for _ , r := range s {
		hash_map[r] += 1
	}

	for _, r := range t {
		hash_map[r] -= 1

		if hash_map[r] < 0 {
			return false
		}
	}

	for _, v := range hash_map {
		if v != 0 {
			return false
		}
	}

	return true
}
