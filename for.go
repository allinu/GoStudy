// Package main provides ...
package main

import "fmt"

func main() {
	// NOTE: 初始语句 判断语句 结束语句
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// var i = 5
	// for ; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// INFO: 99 乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf(" %dx%d=%d ", i, j, i*j)
		}
		fmt.Printf("\n")
	}
}
