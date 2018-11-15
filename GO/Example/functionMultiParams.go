package main

import (
	"fmt"
	"strings"
)

func main() {
	// 标准方式
	fmt.Println(toFullname("wang", "jue"))
	// 切片方式,需要在切片后加上...才行!
	s := []string{"danny", "wang"}
	fmt.Println(toFullname(s...))
}

func toFullname(names ...string) string {
	return strings.Join(names, " ")
}
