package main

import (
	"fmt"
	"newtoallofthis/go_learning/algos/sortings"
	"newtoallofthis/go_learning/utils/calculator"
	"newtoallofthis/go_learning/utils/cli"
)

func main() {
	cmd, nums := cli.Handle_Cli()
	switch cmd {
	case "help":
		fmt.Println("Help")
	case "add":
		fmt.Println("Adding")
		fmt.Println(calculator.Add(nums))
	case "sub":
		fmt.Println("Subtracting")
		fmt.Println(calculator.Sub(nums))
	case "bubble":
		fmt.Println("Bubble")
		fmt.Println(sortings.BubbleSort(nums))
	case "selection":
		fmt.Println("Selection")
		fmt.Println(sortings.SelectionSort(nums))
	case "insertion":
		fmt.Println("Insertion")
		fmt.Println(sortings.InsertionSort(nums))
	default:
		fmt.Println("Invalid command")
	}
}
