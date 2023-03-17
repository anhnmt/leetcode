package find_the_index_of_the_first_occurrence_in_a_string

import (
	"strings"
)

func strStr(haystack string, needle string) int {
	n := len(needle)
	for i := 0; i < len(haystack)-n+1; i++ {
		if haystack[i:i+n] == needle {
			return i
		}
	}
	return -1
}

func strStr1(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
