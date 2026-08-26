type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	encoded_str := ""

	for _, cur := range strs {
		encoded_str += strconv.Itoa(len(cur)) + "#" + cur 
	}

	return encoded_str
}

func (s *Solution) Decode(encoded string) []string {
	i := 0
	result := []string{}

	for i<len(encoded) {
		length := 0
		//process length
		for i < len(encoded) && encoded[i] != '#' {
			cur, _ := strconv.Atoi(string(encoded[i]))
			
			length = length * 10 + cur
			i++
		}

		//skip "#"
		i++

		result = append(result, encoded[i:i+length])
		i += length
	}

	return result
}
