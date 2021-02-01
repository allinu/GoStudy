package main

import "fmt"

// NOTE iota 每增加一行（空行除外）就会加一

// NOTE 下面是常量定义简写

// NOTE _ 黑洞接收

const (
	a = 100
	// NOTE 不写会与上一行值相同
	b
	c
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println(KB, MB)
}
