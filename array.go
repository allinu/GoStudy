// Package main provides ...
package main

import "fmt"

func modifyArray(x [3]int) {
	x[0] = 100
}

func modifyArray2(x [3][2]int) {
	x[2][0] = 100
}

func main() {
	var testArray [10]int
	var numArray = [2]int{1, 2}
	var cityArray = [3]string{"Beijing", "Shanghai", "Guanzhou"}
	fmt.Println(testArray)
	fmt.Println(numArray)
	fmt.Println(cityArray)
	//支持的写法
	a := [...][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	//不支持多维数组的内层使用...
	// b := [3][...]string{
	// 	{"北京", "上海"},
	// 	{"广州", "深圳"},
	// 	{"成都", "重庆"},
	// }
	fmt.Println(a)
	// fmt.Println(b)
	c := [3]int{10, 20, 30}
	modifyArray(c) //在modify中修改的是a的副本x
	fmt.Println(c) //[10 20 30]
	b := [3][2]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	modifyArray2(b) //在modify中修改的是b的副本x
	fmt.Println(b)  //[[1 1] [1 1] [1 1]]

}
