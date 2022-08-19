package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sortArray(nums []int) []int {
	heapSort(nums)
	return nums
}

func heapSort(nums []int) {
	length := len(nums)
	createHeap(nums, length)
	for i := len(nums) - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		// 交换完之后 边界就减一了
		length--
		// 一直是从0 开始往下调整
		shiftDown(nums, 0, length)
	}
}

func createHeap(nums []int, size int) {
	// 找到第一个非叶子节点
	root := (size - 2) / 2
	// 从这里开始向下调整
	for ; root >= 0; root-- {
		shiftDown(nums, root, size)
	}
}

// 从上至下调整
func shiftDown(nums []int, parent, size int) {
	child := 2*parent + 1
	for child < size {
		// 简历大根堆
		if child+1 < size && nums[child] < nums[child+1] {
			child += 1
		}
		if nums[child] > nums[parent] {
			nums[child], nums[parent] = nums[parent], nums[child]
		} else {
			break
		}
		parent = child
		child = 2*parent + 1
	}
}

func main() {
	str := ""
	reader := bufio.NewReader(os.Stdin)
	readString, _ := reader.ReadString('\n')
	str = readString[1 : len(readString)-3]
	arr := strings.Split(str, " ")
	nums := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		nums[i], _ = strconv.Atoi(arr[i])
	}
	sort.Ints(nums)
	str = ""
	for i := 0; i < len(arr); i++ {
		if i == 0 {
			str += "\"" + strconv.Itoa(nums[i]) + " "
		} else if i == len(arr)-1 {
			str += strconv.Itoa(nums[i]) + "\""
		} else {
			str += strconv.Itoa(nums[i]) + " "
		}
	}
	fmt.Println(str)
}
