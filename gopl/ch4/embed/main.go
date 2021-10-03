package main

import "fmt"

//type Circle struct {
//	X, Y, Radius int
//}
//
//type Wheel struct {
//	X, Y, Radius, Spokes int
//}

type Point struct {
	X, Y int
}

//type Circle struct {
//	Center Point
//	Radius int
//}
//
//type Wheel struct {
//	Circle Circle
//	Spokes int
//}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main(){
	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20

	fmt.Println(w)
	fmt.Printf("%#v\n", w)

	w1 := Wheel{Circle{Point{9, 9}, 5}, 20}
	w2 := Wheel{
		Circle: Circle{
			Point:Point{X: 10, Y:10},
			Radius: 15,
		},
		Spokes: 25,
	}

	fmt.Println(w1)
	fmt.Println(w2)
}