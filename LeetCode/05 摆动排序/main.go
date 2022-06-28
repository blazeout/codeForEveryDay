package main

import (
	"fmt"
	"sort"
)

func wiggleSort(nums []int) {
	l := len(nums)
	j, k := l, (l+1)>>1
	t := make([]int, l)
	copy(t, nums)
	sort.Ints(t)
	for i := 0; i < l; i++ {
		if i&1 == 1 {
			j--
			nums[i] = t[j]
		} else {
			k--
			nums[i] = t[k]
		}
	}
}

func main() {
	nums := []int{1, 2, 1, 2, 1, 1, 2, 2, 1}
	wiggleSort(nums)
	fmt.Println(nums)
}
