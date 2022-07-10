package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"sync"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	res := [][]int{}
	for len(queue) > 0 {
		size := len(queue)
		tmp := []int{}
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, tmp)
	}
	return res
}

func GeneratePassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
}

func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

var wg sync.WaitGroup

func getValue(channel chan int) {
	i, ok := <-channel
	fmt.Println(i, ok)
	close(channel)
	i, ok = <-channel
	fmt.Println(i, ok)
	wg.Done()
}

func main() {
	//passWord := "123"
	//passWordErr := "132"
	//hashedPassword, err := GeneratePassword(passWord)
	////generatePassword, _ := GeneratePassword(passWordErr)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//err = ComparePassword(string(hashedPassword), passWord)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("ok1")
	//err = ComparePassword(string(hashedPassword), passWordErr)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("ok2")
	channel := make(chan int)
	go getValue(channel)
	channel <- 1
	wg.Add(1)
	wg.Wait()
}
