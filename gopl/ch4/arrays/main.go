package main

import (
	"fmt"
)

func zero(ptr *[32]byte)  {
	for i := range ptr {
		ptr[i] = 0
	}

}
func main()  {
	//var a [3]int
	//fmt.Println(a[0]) // print the first  element
	//fmt.Println(a[len(a) - 1]) // print the last element
	//
	//// Print the indices and elements
	//for i, v := range a {
	//	fmt.Printf("%d %d\n", i, v)
	//}
	//
	//// Print the elements only
	//for _, v := range a {
	//	fmt.Printf("%d \n", v)
	//}
	//
	//// Array literal
	//var q1 [3]int = [3]int{1, 2, 3}
	//fmt.Println(q1)
	//q2 := [...]int{1, 2, 3, 4}
	//fmt.Printf("%T\n", q2)
	//q3 := [5]int{1, 2}
	//fmt.Println(q3[4])
	//q1 = [...]int{2, 3, 4}
	//fmt.Println(q1)

	//type Currency int
	//const (
	//	USD Currency = iota
	//	EUR
	//	GBP
	//	RMB
	//)
	//symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	//
	//fmt.Println(RMB, symbol[RMB])

	//a := [2]int{1, 2}
	//b := [...]int{1, 2}
	//c := [2]int{1, 3}
	//fmt.Println(a==b, b == c , c == a)
	//d := [3]int{1,2}
	//fmt.Println(a == d)

	//c1 := sha256.Sum256([]byte("x"))
	//c2 := sha256.Sum256([]byte("X"))
	//fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	//
	p := [32]byte{1}
	fmt.Println(p)
	zero(&p)
	fmt.Println(p)
}

