package main

import "fmt"

/*
! CONSTANT GO
 * In Golang : we use the term constant fixed (unchanging) value.
 * We use the keyword const to declare a constant.
 * A constant belongs to the compile time and its created at compile time.
 * We can't change the value of a constant at runtime.

*/

func main() {
	// !declare a constant
	const PI float64 = 3.14 // * if you declear constant in not use code not throw any error

	var i int
	fmt.Println(i) // 0

	x, y := 5, 1 // if i replace 1 to 0

	fmt.Println(x / y) // vs code don't detect error because it run time error

	// const a = 5
	// const b = 0
	// fmt.Println(a / b) // *panic: runtime error: integer divide by zero

	const n, m = 10, 15
	// !define  multipal constant

	const (
		r = 10
		t
		c
	) //! 10 10 10 if i assign one value othe constant is copy to parent value

	fmt.Println(r, t, c)
	// todo : CONSTANT RULE
	//* 1. constant must be initialized
	//* 2. constant must be immutable
	const temp int = 30
	// * you can not initilizied a constant at run time
	// temp = 40
	// s := 5
	// const tt = s // this is error

	// const x int = 5
	// const y float64 = 4.5 * x // this is error

	const xx = 5
	const yy = 4.5 * 5 // this is correct

	fmt.Printf("%T\n", yy)

	const rr = 5
	var rt = rr
	fmt.Printf("%T\n", rt)

}
