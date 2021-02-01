package main

import "fmt"

func main() {
	i1 := 101
	// OUTPUT 1100101
	fmt.Printf("%b\n", i1)
	// OUTPUT 101
	fmt.Printf("%d\n", i1)
	// OUTPUT 145
	fmt.Printf("%o\n", i1)
	// OUTPUT 65
	fmt.Printf("%x\n", i1)
	i2 := 077
	// OUTPUT 63
	fmt.Printf("%d\n", i2)
	i3 := 0x2e3
	// OUTPUT 739
	fmt.Printf("%d\n", i3)
	// OUTPUT int
	fmt.Printf("%T\n", i3)
}
