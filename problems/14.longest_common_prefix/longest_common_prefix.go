package longest_common_prefix

func longestCommonPrefix(strs []string) (pre string) {
	pre = strs[0]
	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(pre); j++ {
			if len(strs[i]) <= j || strs[i][j] != pre[j] {
				pre = pre[0:j]
				break
			}
		}
	}

	return
}
