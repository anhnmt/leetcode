package valid_parentheses

func isValid(s string) bool {
	// Kiểm tra chuỗi rỗng
	if len(s) == 0 {
		return true
	}

	stack := make([]rune, 0)
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
			continue
		}

		l := len(stack)

		// Kiểm tra các điều kiện
		// nếu v == ")" || "]" || "}"
		// và l > 0
		// và giá trị cuối cùng trong stack == "(" || "[" || "{"
		// thì sẽ xoá stack từ đầu cho tới l-1
		if (v == ')' && l > 0 && stack[l-1] == '(') ||
			(v == ']' && l > 0 && stack[l-1] == '[') ||
			(v == '}' && l > 0 && stack[l-1] == '{') {
			stack = stack[:l-1]
		} else {
			return false
		}
	}

	// Kiểm tra stack rỗng
	return len(stack) == 0
}
