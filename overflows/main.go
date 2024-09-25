package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x uint8 = 255
	x++ // overflow , xis 0

	fmt.Println(x)

	// a := int8(255 + 1)
	// fmt.Println(a)// Constant 256 overflows int8

	var t = 3   // ! type int
	var u = 3.5 // * type float64

	y := t * int(u)

	fmt.Println(y)

	fmt.Printf("%T\n", t)
	fmt.Printf("%T\n", u)

	// int and int64 not a equal
	var a = 5
	var b int64 = 5

	// fmt.Println(a == b) //* invalid operation: a == b (mismatched types int and int64)compilerMismatchedTypes)
	fmt.Println(a == int(b))

	s := string(99)
	fmt.Println(s) //c

	//s1 := string(66.5) //*cannot convert 66.5 (untyped float constant) to type stringcompilerInvalidConversion
	var mystr = fmt.Sprintf("%f\n", 44.2)
	fmt.Println("value of mystr is ", mystr)
	fmt.Printf("%T\n", mystr)

	var mystr1 = fmt.Sprintf("%d\n", 34234)
	fmt.Println("value of mystr1 is ", mystr1)
	fmt.Printf("%T\n", mystr1)
	fmt.Println(string(35234))

	s1 := "3.565"
	fmt.Printf("%T\n", s1) //string

	var f1, err = strconv.ParseFloat(s1, 64)

	_ = err
	fmt.Println(f1)
	/*
			!func strconv.ParseFloat(s string, bitSize int) (float64, error)
			*ParseFloat converts the string s to a floating-point number with the precision specified by bitSize: 32 for float32, or 64 for float64.
		   * When bitSize=32, the result still has type float64, but it will be convertible to float32 without changing its value
	*/

	i, err := strconv.Atoi("-50") //I type is int, i value is -50
	s2 := strconv.Itoa(20)        //S2 type is string, S2 value is 20

	fmt.Printf("I type is %T, i value is %v\n", i, i)
	fmt.Printf("S2 type is %T, S2 value is %v\n", s2, s2)

	// ! math.Pow() to calculate the power

}
