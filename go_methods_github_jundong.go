package main

import "fmt"

//go 语言字符串反转功能,双指针
func reverse(str string) string {
	str1 := []byte(str)
	for i, j := 0, len(str) - 1; i < j; i, j = i + 1, j - 1 {
		fmt.Println(i,j)
		str1[i], str1[j] = str1[j], str1[i]
	}
	return string(str1)
}