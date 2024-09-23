package main

import (
	"fmt"
)

/*
	! fmt package in details
	* the varbs inside the template string start with percent sign followed by a single letter
	?Exampe
	! %d is a verb tells Printf that it should replace it with an integer.
	* duble cort inside duble cort go is not allow slash then you can use
*/

func main() {
	fmt.Println("Hello, World!")

	fmt.Printf("your %d is your age\n", 21)

	fmt.Printf("X is %d, y is %f", 5, 7.5)
	//fmt.Printf("He says: "Hello Go!"")// this is not allow
	fmt.Printf("he says: \"hello!\"") // it allow

	figure := "circle"
	radius := 5
	pi := 3.14159

	_, _ = figure, pi
	fmt.Printf("Redius is %+d\n ", radius)

	fmt.Printf("this is PI value %f\n", pi)

	// ! %q for quoted string
	fmt.Printf("this is a %q\n", figure)

	// ! %v -> replace by any type

	fmt.Printf("this diameter of a %v with a radius of %v is %v\n", figure, radius, pi)

	// ! %T -> type
	fmt.Printf("figure is of type  %T\n", figure) // string
	fmt.Printf("radius is of type %T\n", radius)  // Int
	fmt.Printf("Pi is of type %T\n", pi)          // float64

	// ! %t -> bool
	close := true
	fmt.Printf("file  close %t\n", close)

	// ! %b = base 2
	fmt.Printf("binary of 5 is %08b\n", 55)
	fmt.Printf("Decimal of 100 is %x\n", 100)

	x := 5.3
	y := 6.7

	fmt.Printf("X * Y = is %.3f\n", x*y)
}
