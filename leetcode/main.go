package main

import (
	"fmt"

	"github.com/becas/go_prac/leetcode/two_sum"
)

func main() {
	// 1. two sum
	twoSum()
}

// creating so to
func twoSum() {
	arr := []int{1, 3, 5, 7, 8, 2}
	target := 9

	result := two_sum.TwoSumBruteForce(arr, target)
	fmt.Println(result)

	result1 := two_sum.TwoSumHash(arr, target)
	fmt.Println(result1)
}
