package cli

import (
	"fmt"
	"os"
	"strconv"
)

func Get_Args() []string {
	return os.Args[1:]
}

func Return_Action() string {
	args := Get_Args()
	if len(args) == 0 {
		return "help"
	}
	return args[0]
}

func Handle_Cli() (string, []int) {
	args := Get_Args()
	var nums []int
	for _, arg := range args[1:] {
		intValue, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("Error converting string to int: %v\n", err)
			os.Exit(1)
		}
		nums = append(nums, intValue)
	}
	return args[0], nums
}
