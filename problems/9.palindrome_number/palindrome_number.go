package palindrome_number

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	if x%10 == 0 {
		return false
	}

	a := make([]int, 0, 32)
	for x > 0 {
		a = append(a, x%10)
		x = x / 10
	}

	l := len(a)
	for i, j := 0, l-1; i <= j; i, j = i+1, j-1 {
		if a[i] != a[j] {
			return false
		}
	}

	return true
}
