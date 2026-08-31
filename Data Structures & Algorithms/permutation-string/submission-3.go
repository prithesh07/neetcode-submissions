func checkInclusion(s1 string, s2 string) bool {
	left := 0
	s1Len := len(s1)
	s2Len := len(s2)

	if s2Len < s1Len {
		return false
	}

	s1Window := make([]int, 26)
	s2Window := make([]int, 26)

	for i := 0; i < s1Len; i++ {
		s1Window[s1[i] - 'a'] += 1
	}

	for i := 0; i < s1Len; i++ {
		s2Window[s2[i] - 'a'] += 1
	}

	matches := 0

	for i:=0 ; i<26; i++ {
		if s1Window[i] == s2Window[i] {
			matches++
		}
	}

	for {
		if matches == 26 {
			return true
		}
		
		s2Left := s2[left]
		
		if s2Window[s2Left - 'a'] == s1Window[s2Left - 'a'] {
			s2Window[s2Left - 'a'] -= 1
			matches--
		} else {
			s2Window[s2Left - 'a'] -= 1
			if s2Window[s2Left - 'a'] == s1Window[s2Left - 'a'] {
				matches++
			}
		}

		left++

		if left + s1Len - 1 >= s2Len {
			break
		}

		s2Right := s2[left + s1Len - 1]

		if s2Window[s2Right - 'a'] == s1Window[s2Right - 'a'] {
			matches--
			s2Window[s2Right - 'a'] += 1
		} else {
			s2Window[s2Right - 'a'] += 1

			if s2Window[s2Right - 'a'] == s1Window[s2Right - 'a'] {
				matches++
			}
		}
	}

	return false
}
