package main

import (
	"fmt"
	"math/rand"
)

// rand5 实现 rand7

func rand5() int {
	return rand.Intn(5) + 1
}

func rand7() int {
	for {
		row := rand5()
		col := rand5()
		indexer := (row-1)*5 + col
		if indexer <= 20 {
			return (indexer-1)%7 + 1
		}
	}
}

func rand10() int {
	for {
		row := rand7()
		col := rand7()
		indexer := (row-1)*7 + col
		if indexer <= 40 {
			return (indexer-1)%10 + 1
		}
	}
}

func main() {
	for i := 0; i < 100; i++ {
		i2 := rand10()
		fmt.Print(i2, " ")
	}
}
