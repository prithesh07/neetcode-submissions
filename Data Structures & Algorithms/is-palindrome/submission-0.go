func isPalindrome(s string) bool {
	leftPtr := 0 
	rightPtr := len(s) - 1

	for leftPtr < rightPtr {
		for !unicode.IsLetter(rune(s[leftPtr])) && !unicode.IsDigit(rune(s[leftPtr])) {
			leftPtr++

			if leftPtr > rightPtr {
				return true
			}
		}

		for !unicode.IsLetter(rune(s[rightPtr])) && !unicode.IsDigit(rune(s[rightPtr])) {
			rightPtr--

			if leftPtr > rightPtr {
				return true
			}
		}

		if unicode.ToLower(rune(s[leftPtr])) != unicode.ToLower(rune(s[rightPtr])) {
			return false
		}

		leftPtr++
		rightPtr--
	}

	return true
}
