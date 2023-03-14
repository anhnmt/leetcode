package valid_parentheses

import (
	"problems/structures"
)

func isValid(s string) bool {
	// Khởi tạo stack
	l := len(s)
	st := structures.NewStack(l)
	for i := 0; i < l; i++ {
		// Nếu s[i] == '(' thì thêm thẻ đóng ')' vào stack
		if s[i] == '(' {
			st.Push(')')
			continue
		}

		// Nếu s[i] == '[' thì thêm thẻ đóng ']' vào stack
		if s[i] == '[' {
			st.Push(']')
			continue
		}

		// Nếu s[i] == '{' thì thêm thẻ đóng '}' vào stack
		if s[i] == '{' {
			st.Push('}')
			continue
		}

		// Kiểm tra nếu s[i] == ')' || ']' || '}'
		// Thì lấy Pop trong stack và so sánh với s[i]
		// nếu không thì trả về false
		if s[i] == ')' || s[i] == ']' || s[i] == '}' {
			if st.Pop() != s[i] {
				return false
			}
		}
	}

	// Kiểm tra stack rỗng
	return st.IsEmpty()
}
