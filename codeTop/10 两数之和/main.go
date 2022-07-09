package _0_两数之和

func twoSum(nums []int, target int) []int {
	hashMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if j, OK := hashMap[target-nums[i]]; OK {
			return []int{i, j}
		} else {
			hashMap[nums[i]] = i
		}
	}
	return nil
}
