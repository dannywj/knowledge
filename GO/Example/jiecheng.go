package main

//  3!=4*3*2*1
import "fmt"

func main() {
	fmt.Println("======begin=====")
	for i := 1; i <= 10; i++ {
		result := compute(i)
		output := fmt.Sprintf("%d-result:%d", i, result)
		fmt.Println(output)
	}

}

func compute(n int) int {
	if n > 1 {
		return n * compute(n-1)
	} else {
		return 1
	}
}