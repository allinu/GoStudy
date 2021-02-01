package main

import "fmt"

func main() {
	b1 := true
	// NOTE: 默认值是false
	var b2 bool
	fmt.Printf("%T\t", b1)
	// OUTPUT: false
	fmt.Printf("%v\n", b2)
}
