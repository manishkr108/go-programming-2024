package main

import "fmt"

/*
* ------------Define Type ------------
? a defined type also called a named type is a new type created by the programmer from another existing type which is called the underlying
? or source type
? a defined type is a new name given to an existing type


*/

func main() {
	// type age int        // * int is its undelying type
	// type oldAge age     // * int is  its underlying type
	// type varyoldAge age //* int is  its underlying type

	type speed uint

	var s1 speed = 10
	var s2 speed = 20

	fmt.Println(s2 - s1)

	var x uint
	// x = s1  // * error: deffrent type
	x = uint(s1)
	_ = x
	var s3 speed = speed(x)
	fmt.Println(s3)

	//! ---------Alias Declaration----------------

	var a uint8 = 10
	var b byte
	b = a
	_ = b

	type second = uint
	var hour second = 3600
	fmt.Printf("minutes in an hour : %d\n", hour/60)

}
