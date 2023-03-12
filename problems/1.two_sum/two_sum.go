package two_sum

func twoSum(nums []int, target int) []int {
	// tạo 1 map kiểu int
	m := make(map[int]int)

	// chạy vòng lặp để lấy giá trị trong map
	for i, num := range nums {
		// kiểm tra xem target - num đã tồn tại trong map chưa
		// nếu có thì trả về index trong map và i
		if idx, ok := m[target-num]; ok {
			return []int{idx, i}
		}

		// lưu lại vị trí của num vào map
		m[num] = i
	}

	return nil
}
