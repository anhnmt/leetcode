package length_of_last_word

func lengthOfLastWord(s string) int {
	c := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			c++
			continue
		}

		if c > 0 {
			break
		}
	}

	return c
}
