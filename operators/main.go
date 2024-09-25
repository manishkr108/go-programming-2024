package main

import "fmt"

/*
! -----------OPERATORS--------
* An operators is a symbol of the programming language which is able to "operate on value".
* Operators are used to perform operations on variables and values.
	! 1. Arthmetic and BitWise Operators : +,-,*,/,%,&,|,^,<<,>>
	! 2. Assignment Operators : =,+=,-=,*=,/=
	! 3. Compairsion Operators: ==,!=, <,>, >=,
	! 4. Logical Operators : &&,||,!
	! 5. Operators for Pointers (&) and channels (<-)

*/

func main() {
	// Arithmetic

	a, b := 6, 2
	r := (a + b) / (a - b) * 2
	fmt.Println(r) // 6

	c, d := 2, 3
	c += d
	fmt.Println(c)

	x := 3
	x += 1
	x++
	fmt.Println(x) //5
	// fmt.Println(5 + x++) // this is incorrect way in go
	fmt.Println(a == b) // false
	fmt.Println(a != b) // true
	fmt.Println(a, b)
	fmt.Println(a > b)  // true
	fmt.Println(a < b)  // false
	fmt.Println(a >= b) // true
	fmt.Println(a <= b) // false
}
