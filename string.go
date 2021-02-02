package main

import (
	"fmt"
	"strings"
)

func main() {
	tmp := "\"C:\\windows\\\""
	fmt.Printf(tmp + "\n")
	// NOTE: 多行的字符串
	tmp_1 := `
		世情薄
		人情恶
		雨送黄昏花易落
	`
	fmt.Printf(tmp_1)
	fmt.Println(len(tmp))
	fmt.Println(tmp + tmp)
	ss := fmt.Sprintf("%s%s", tmp, tmp)
	fmt.Printf(ss)

	ret := strings.Split(tmp, "\\")
	fmt.Println(ret)
	fmt.Println(ret[0])
	fmt.Println(strings.Join(ret, "+"))

	s1 := "白萝卜"
	s2 := []rune(s1)
	s2[0] = '胡'
	fmt.Println(string(s2))
	var count int
	// INFO: count the chinese symbol
	in_str := "Hello白雪公主和七个小矮人"
	for _, i := range in_str {
		if i > 255 {
			count++
			fmt.Printf("%c", i)
		}
	}
	// OUTPUT: 10
	fmt.Println(count)
}
