package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	var numbers [5]int
	fmt.Printf("%v\n", numbers)  //[0 0 0 0 0]
	fmt.Printf("%#v\n", numbers) //[5]int{0, 0, 0, 0, 0}

	// other way declear array

	var a1 = [4]float64{}
	fmt.Printf("%v\n", a1) //[0 0 0 0]

	var a2 = [3]int{-10, 1, 100}
	fmt.Printf("%v\n", a2) //[-10 1 100]

	//short declaration
	a3 := [4]string{"manish", "kumar", "raj", "go lang"}
	fmt.Printf("%#v\n", a3) //[4]string{"manish", "kumar", "raj", "go lang"}

	a4 := [5]string{"x", "y"}
	fmt.Printf("%#v\n", a4) //[5]string{"x", "y", "", "", ""}

	// a nice feature of go is the ellipsis operator or triple dots. it finds out automatically the length of array
	a5 := [...]int{1, 2, 3, 4, 5, 5, 6, 7, 8, 89, 7}
	fmt.Printf("%#v\n", a5)                         //[11]int{1, 2, 3, 4, 5, 5, 6, 7, 8, 89, 7}
	fmt.Printf("the length of a5 is %d\n", len(a5)) //the length of a5 is 11
	fmt.Printf("%T\n", a5)

	// Use reflect to get the type
	t := reflect.TypeOf(a5)

	// Check if it's an array or slice
	if t.Kind() == reflect.Array {
		fmt.Println("This is an array")
	} else if t.Kind() == reflect.Slice {
		fmt.Println("This is a slice")
	}

	/*
		In Go, Kind() is a method provided by the reflect package that returns the specific kind of a value.
		 It helps you determine the underlying data type of an object, whether it's a struct, array, slice, map, pointer, etc.
	*/

	abc := [...]int{2, 3, 4, 5, 6, 7, 0}
	fmt.Printf("%#v\n", abc) //[7]int{2, 3, 4, 5, 6, 7, 0}
	fmt.Printf("%v\n", abc)  // [2 3 4 5 6 7 0]

	a6 := [...]int{1,
		2,
		3,
		4,
		5, // in last element must add comma
	}
	fmt.Println(a6)

	//  Add and remove elements

	a6[0] = 9
	fmt.Println(a6) //[9 2 3 4 5]

	// a6[10] =12 // invalid argument: index 10 out of bounds -> there is no element index 10

	//How to Iterate array (how many way)
	// 1. using range
	//range = here range is just a language key word used for iteration.
	for i, v := range a6 {
		fmt.Println("Index: ", i, "Value:", v)
	}

	fmt.Println(strings.Repeat("#", 20))
	//2nd  way

	for i := 0; i < len(a6); i++ {
		fmt.Println("Index: ", i, "Value:", a6[i])
	}

	// Lets create multi dimention array

	balance := [2][3]int{
		{1000, 2000, 3000},
		{4000, 5000, 6000},
	}

	fmt.Println(balance)

	m := [3]int{1, 2, 3}

	n := m

	fmt.Println("n equal to m :", n == m) //n equal to m : true
	m[0] = -1

	fmt.Println("n m :", n, m) //n m : [1 2 3] [-1 2 3]

	fmt.Println("n equal to m :", n == m) //n equal to m : false
	// 2 arrays are equal if they have the same length and the same elements, in the same order

	// Array with keyed Elements

	grade := [3]int{
		1: 90,
		0: 70,
		2: 80,
	}

	fmt.Println(grade) // [70 90 80]

	acco := [3]int{2: 50}
	fmt.Println(acco) //[0 0 50]

	names := [...]string{
		5: "paris",
	}
	fmt.Println(names, len(names)) //[     paris] 6

	city := [...]string{
		5:       "delhi",
		"Noida", // index 6 insert in last if key not there
		1:       "India",
	}

	fmt.Printf("%#v\n", city) //[7]string{"", "India", "", "", "", "delhi", "Noida"}

	weekend := [7]bool{5: true, 6: true}
	fmt.Println(weekend) // [false false false false false true true]
}
