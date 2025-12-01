package main

import (
	"fmt"

	"github.com/becas/go_prac/dsa/binary_search"
)

func main() {
	arr := []int{1, 2, 3, 5, 6, 7, 8, 9, 10}
	num := 10
	index := binary_search.Binary_search(arr, num)
	if index >= 0 {
		fmt.Println("Found the numeber", num)
	} else {
		fmt.Println("Helppp we lost it ")
	}
}
