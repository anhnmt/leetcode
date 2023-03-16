package plus_one

func plusOne(digits []int) []int {
	// chạy từ dưới lên đầu
	// nếu digits[i] + 1 < 9, thì lấy digits[i] + 1 và return digits
	// ngược lại thì gán digits[i] = 0
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	// thêm 1 vào đầu mảng
	digits = append([]int{1}, digits...)
	return digits
}
