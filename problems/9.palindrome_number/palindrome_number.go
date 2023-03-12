package palindrome_number

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var a []int
	for x != 0 {
		a = append(a, x%10)
		x = x / 10
	}

	for i := 0; i < len(a); i++ {
		if a[i] != a[len(a)-i-1] {
			return false
		}
	}

	return true
}
