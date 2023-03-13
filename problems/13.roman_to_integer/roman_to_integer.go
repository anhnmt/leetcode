package roman_to_integer

var m = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) (t int) {
	l, n, lastint := len(s), 0, 0
	for i := 0; i < l; i++ {
		char := s[l-(i+1) : l-i]
		n = m[char]
		if n < lastint {
			t -= n
		} else {
			t += n
		}

		lastint = n
	}

	return
}
