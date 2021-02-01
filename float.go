package main

import (
	"fmt"
	"math"
)

func main() {
	// OUTPUT: 3.4028234663852886e+38
	fmt.Println(math.MaxFloat32)
	f1 := 1.23456
	// NOTE: 默认的是float64，如果使用float32需要强制转换
	// NOTE: float32 的值不能直接赋值float64
	fmt.Printf("%T\n", f1)
}
