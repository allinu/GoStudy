package main

import "fmt"

// 函数外只能声明变量

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7}
	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Println("Hello World!")
}
