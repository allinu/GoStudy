// Package main provides ...
package main

import "fmt"

func main() {
	n := 100
	fmt.Printf("%T\n", n)
	fmt.Printf("%#v\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)
	name := "Tony"
	fmt.Printf("Hello %s\n", name)
	// NOTE: %#V 会添加类型描述符
	fmt.Printf("Hello %#v\n", name)
	// NOTE: %#V 会添加类型描述符
	list_demo := [...]int{1, 2, 3}
	fmt.Printf("%v => ", list_demo)
	fmt.Printf("%#v\n", list_demo)
}
