package main

/*
Go读取大文件-按行读取
*/
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("========begin read=========")
	filename := "redbag.txt"
	ReadLine(filename)
	fmt.Println("========end read=========")
}

func ReadLine(filename string) {
	f, _ := os.Open(filename)
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		str, err := readLine(r)
		if err != nil {
			break
		}
		fmt.Println(str)
	}
}

func readLine(r *bufio.Reader) (string, error) {
	line, isprefix, err := r.ReadLine()
	for isprefix && err == nil {
		var bs []byte
		bs, isprefix, err = r.ReadLine()
		line = append(line, bs...)
	}
	return string(line), err
}
