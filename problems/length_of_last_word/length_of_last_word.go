package length_of_last_word

func lengthOfLastWord(s string) int {
	c := 0
	for i := len(s) - 1; i >= 0; i-- {
		if c > 0 && s[i] == ' ' {
			return c
		}

		if s[i] == ' ' && c == 0 {
			s = s[:i]
			continue
		}

		c++
	}

	return len(s)
}
