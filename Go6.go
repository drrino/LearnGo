package main

import "fmt"

/**
 * go没有三目运算符，需使用完整的if判断
 */
func main()  {
	if 7 % 2==0 {
		fmt.Println("7 is even")
	}else{
		fmt.Println("7 is odd")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}