package _2_比较版本号

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	arr1, arr2 := strings.Split(version1, "."), strings.Split(version2, ".")
	for i := 0; i < len(arr1) || i < len(arr2); i++ {
		temp2 := 0
		temp1 := 0
		if i < len(arr2) {
			temp2, _ = strconv.Atoi(arr2[i])
		}
		if i < len(arr1) {
			temp1, _ = strconv.Atoi(arr1[i])
		}
		if temp1 > temp2 {
			return 1
		} else if temp1 < temp2 {
			return -1
		}
	}
	return 0
}
