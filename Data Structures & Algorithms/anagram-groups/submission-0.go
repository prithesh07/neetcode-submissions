func getKey(str string) string {
	counter := make([]int, 26)

	for _, char := range str {
		counter[int(char)-97] += 1
	}

	data, _ := json.Marshal(counter)

	return string(data)
}

func groupAnagrams(strs []string) [][]string {
	hash_map := make(map[string][]string)

	for _, val := range strs {
		key := getKey(val)

		list, exists := hash_map[key]

		if exists {
			hash_map[key] = append(list, val)
		} else {
			hash_map[key] = []string {val}
		}
	}

	result := make([][]string, 0, len(hash_map))

	for _, val := range hash_map {
		result = append(result, val)
	}

	return result
}
