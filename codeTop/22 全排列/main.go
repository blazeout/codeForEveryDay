package _2_全排列

func permute(nums []int) [][]int {
	mp := map[int]bool{}
	var dfs func(path []int, index int)
	ans := [][]int{}
	dfs = func(path []int, index int) {
		// 谨记: 不能直接把path加入到ans中, 因为path是一个引用, 如果直接把path加入到ans中, 后续改的则会导致path被改变
		if index == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if mp[i] {
				continue
			}
			mp[i] = true
			path = append(path, nums[i])
			dfs(path, index+1)
			mp[i] = false
			path = path[:len(path)-1]
		}
	}
	dfs([]int{}, 0)
	return ans
}
