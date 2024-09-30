package main

import (
	"fmt"
	"reflect"
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

}
