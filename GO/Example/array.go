package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("==========begin==========")
	//array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	slice := []int{1, 2, 10, 4, 5, 6, 7, 8, 9, 3}
	fmt.Println("init array valus:")
	fmt.Println(slice)
	fmt.Println("begin get top 3")

	var re = getTopThreeValues(slice)
	fmt.Println("result")
	fmt.Println(re)
}

func getTopThreeValues(slice_data []int) []int {
	var max_val = slice_data[0]
	for _, value := range slice_data {
		fmt.Println(value)
		if value > max_val {
			max_val = value
		}
	}

	fmt.Println("max val:" + strconv.Itoa(max_val))
	news := slice_data[1:3]
	return news
}
