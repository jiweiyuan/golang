package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main()  {
	month := [...]string{
		1: "January",
		2: "February",
		3: "March",
		4: "April",
		5: "May",
		6: "June",
		7: "July",
		8: "August",
		9: "September",
		10: "October",
		11: "November",
		12: "December",
	}

	Q2 := month[4:7]
	summer := month[6:9]
	fmt.Println(Q2, summer)

	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}

	//fmt.Printf("Length of month: %d\n", len(month))
	//fmt.Printf("Capacity of month: %d\n", cap(month))
	//fmt.Printf("Capacity of Q2: %d\n", cap(Q2))
	//
	////fmt.Println(summer[:20])
	//endlessSummer := summer[:5]
	//fmt.Println(endlessSummer)
	//
	//a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//reverse(a[:])
	//fmt.Println(a)

	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)
}
