package _8_合并区间

import "sort"

// 排序后再进行合并即可

func merge(intervals [][]int) [][]int {
	res := [][]int{}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	first, second := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		// 分为有重合区域和没有重合区域两种情况呗~
		if intervals[i][0] <= second {
			// 有重合不合并
			second = max(second, intervals[i][1])
		} else {
			// 没有重合就合并并更新first 和 second
			res = append(res, []int{first, second})
			first = intervals[i][0]
			second = intervals[i][1]
		}
	}
	res = append(res, []int{first, second})
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
