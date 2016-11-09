// getmax.go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	println("======begin get max value from array======")
	// 定义数组
	num_list := []int{1, 3, 5, 7, 9, 6, 8}
	printArray(num_list)
	var max int
	max = getMax(num_list)
	fmt.Println("-- the max value is: [" + strconv.Itoa(max) + "]")
	fmt.Println("-- begin update array")
	updateMaxTo100(&num_list, max)
	fmt.Println("-- finish update array")
	printArray(num_list)
	println("======end get max value from array======")
}

// 数组值传递
func getMax(n_list []int) int {
	var max int
	max = n_list[0]
	for _, val := range n_list {
		if val > max {
			max = val
		}
	}
	return max
}

// 数组引用传递
func updateMaxTo100(n_list *[]int, max int) {
	for index, val := range *n_list {
		if val == max {
			(*n_list)[index] = 100
		}
	}
}

func printArray(n_list []int) {
	fmt.Println("print result:")
	for index, val := range n_list {
		fmt.Println(strconv.Itoa(index+1) + " => " + strconv.Itoa(val))
	}
}
