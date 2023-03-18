package single_number

func singleNumber(nums []int) (res int) {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}

	for _, n := range nums {
		res ^= n
	}
	return res
}
