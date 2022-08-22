package _7_得到目标值的最少行动次数

func minMoves(target int, maxDoubles int) int {
	if maxDoubles == 0 {
		return target - 1
	}
	count := 0
	for target > 1 {
		if target%2 == 1 {
			// 此时就减一
			target -= 1
			count++
		} else if maxDoubles > 0 {
			target /= 2
			count++
			maxDoubles--
		} else if maxDoubles == 0 {
			count += target - 1
			target = 1
		}
	}
	return count
}
