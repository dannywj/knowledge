package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("======begin=====")
	ReadLine("test.txt")
	fmt.Println("======end=====")
}

func checkData(log string) {
	if strings.Contains(strings.ToLower(log), "error") {
		fmt.Println(log)
	}
}

func ReadLine(filename string) {
	f, _ := os.Open(filename)
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		data, err := readLine(r)
		checkData(data)
		if err != nil {
			break
		}
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

// func ReadBlock(filePth string, bufSize int, hookfn func([]byte)) error {
// 	f, err := os.Open(filePth)
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()

// 	buf := make([]byte, bufSize) //一次读取多少个字节
// 	bfRd := bufio.NewReader(f)
// 	for {
// 		n, err := bfRd.Read(buf)
// 		hookfn(buf[:n]) // n 是成功读取字节数

// 		if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
// 			if err == io.EOF {
// 				return nil
// 			}
// 			return err
// 		}
// 	}

// 	return nil
// }
