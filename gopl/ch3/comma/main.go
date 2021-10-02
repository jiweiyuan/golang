package main

import "fmt"

func comma(s string) string  {
	n := len(s)

	if n <= 3 {
		return s
	}
	// recursive, very beautiful
	return comma(s[:n-3]) + "," + s[n-3:]
}

func main()  {
	fmt.Println("Hello World!")
	fmt.Println(comma("1234567"))
	fmt.Println(comma("abcdefg"))
	fmt.Println(comma("中国电信系统集成"))
}