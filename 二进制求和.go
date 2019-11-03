package main

import (
	"fmt"
	"math"
	"strconv"
)


//给定两个二进制字符串，返回他们的和（用二进制表示）。
//
//输入为非空字符串且只包含数字 1 和 0。
//
//示例 1:
//
//输入: a = "11", b = "11"
//输出: "100"
//示例 2:
//
//输入: a = "1010", b = "1011"
//输出: "10101"
//111
func main()  {
	a := "11"
	b := "1"
	num := addBinary(a ,b)
	fmt.Println(num)
	//TenTransTwo(num)
	//convertToBin(6)
	test()
}
func addBinarys(a string, b string) string {
	var a_num float64
	var b_num float64

	for i:=0 ; i < len(a) ; i++{
		mi := len(a) - i -1
		nums_a ,_ := strconv.Atoi(string(a[i]))
		//fmt.Println(nums_a)
		a_num += float64(nums_a)  * math.Pow(float64(2),float64(mi))
	}

	for i:=0 ; i < len(b) ; i++{
		mi := len(b) - i -1
		nums_b ,_ := strconv.Atoi(string(b[i]))
		b_num +=  float64(nums_b)  * math.Pow(float64(2),float64(mi))
	}
	fmt.Println(a_num,b_num)
	number := a_num + b_num
	s := ""
	if number == 0 {
		return "0"
	}
	strNumbers := strconv.FormatFloat(number, 'g', -1, 64)
	fmt.Println(strNumbers)
	intNumbers,err := strconv.Atoi(strNumbers)
	if err!= nil{
		var newNum float64
		intNumbers, err := fmt.Sscanf(strNumbers, "%e", &newNum)
		if err == nil{
			for ; intNumbers > 0 ; intNumbers /= 2 {
				twoNums :=intNumbers % 2
				//fmt.Println(twoNums)
				s = strconv.FormatInt(int64(twoNums),10) + s
			}
			return s
		}
		fmt.Println(intNumbers)
	}


	if err ==nil{
		for ; intNumbers > 0 ; intNumbers /= 2 {
			twoNums :=intNumbers % 2
			//fmt.Println(twoNums)
			s = strconv.FormatInt(int64(twoNums),10) + s
		}

	}
	fmt.Println(s)
	return s
}



func TenTransTwo( number float64) string{
	s := ""
	if number == 0 {
		return "0"
	}
	strNumbers := strconv.FormatFloat(number, 'g', -1, 64)
	fmt.Println(strNumbers)
	intNumbers,err := strconv.Atoi(strNumbers)
	if err!= nil{
		var newNum float64
		intNumbers, err := fmt.Sscanf(strNumbers, "%e", &newNum)
		if err == nil{
			for ; intNumbers > 0 ; intNumbers /= 2 {
				twoNums :=intNumbers % 2
				//fmt.Println(twoNums)
				s = strconv.FormatInt(int64(twoNums),10) + s
			}
			return s
		}
		fmt.Println(intNumbers)
	}


	if err ==nil{
		for ; intNumbers > 0 ; intNumbers /= 2 {
			twoNums :=intNumbers % 2
			//fmt.Println(twoNums)
			s = strconv.FormatInt(int64(twoNums),10) + s
		}

	}
	fmt.Println(s)
	return s
}

func test()  {
	//s += "3"
	f := 1123213213.0
	// 如果格式标记为 'e'，'E'和'f'，则 prec 表示小数点后的数字位数
	// 如果格式标记为 'g'，'G'，则 prec 表示总的数字位数（整数部分+小数部分）
	fmt.Println(strconv.FormatFloat(f, 'f', -1, 32))
	fmt.Println(strconv.FormatFloat(f, 'G', -1, 32))


}
