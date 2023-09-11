package main

import (
	"fmt"
	"newtoallofthis/go_learning/algos/sortings"
	"newtoallofthis/go_learning/utils/cli"
	"strconv"
)

func main() {
	args := cli.Get_Args()
	var nums []int
	for _, arg := range args {
		intValue, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("Error converting string to int: %v\n", err)
			return
		}
		// Append the integer to the integer array
		nums = append(nums, intValue)
	}
	fmt.Println(sortings.BubbleSort(nums))
}
