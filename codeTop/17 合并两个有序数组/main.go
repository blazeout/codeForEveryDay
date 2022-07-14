package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	indexOfNums1 := m - 1
	indexOfNums2 := n - 1
	tail := m + n - 1
	for tail >= 0 {
		if indexOfNums1 >= 0 && indexOfNums2 >= 0 {
			if nums1[indexOfNums1] >= nums2[indexOfNums2] {
				nums1[tail] = nums1[indexOfNums1]
				indexOfNums1--
			} else {
				nums1[tail] = nums2[indexOfNums2]
				indexOfNums2--
			}
		} else if indexOfNums1 < 0 {
			nums1[tail] = nums2[indexOfNums2]
			indexOfNums2--
		} else {
			nums1[tail] = nums1[indexOfNums1]
			indexOfNums1--
		}
		tail--
	}
	return
}

func main() {
	nums1 := []int{1, 3, 6, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	println(nums1)
}
