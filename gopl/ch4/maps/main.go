package main

import (
	"fmt"
	"sort"
)

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, v := range x {
		if yv, ok := y[k]; !ok || yv != v {
			return  false
		}
	}

	return true
}

func main() {
	ages1 := make(map[string]int)
	ages1["alice"] = 31
	ages1["charlie"] = 34
	ages1["blog"] = 28
	fmt.Println(ages1)

	ages2 := map[string]int{
		"alice": 31,
		"charlie": 34,
	}
	delete(ages2, "alice")
	fmt.Println(ages2)
	ages2["blog"] = ages2["blog"] + 1
	fmt.Println(ages2["blog"])

	for name, age := range ages1 {
		fmt.Printf("%s\t%d\n", name, age)
	}

	var names []string

	for name := range ages1 {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages1[name])
	}

	fmt.Println(equal(ages1, ages2))
	fmt.Println(equal(ages2, ages2))


}
