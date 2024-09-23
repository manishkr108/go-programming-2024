package main

import "fmt"

/*
todo - A variable is a name for a memory location where a value od a spacific type is stored
! in Go a varibale belongs and its created at runtime.
? A decleared variable must be used or we get an error
! _ is a blank identifier and mutes the compile time error returend variables.

todo ------DECLEARING VARIABLE-------------
? var x int =7
! var s1 string
? s1 = "Learning Go"

todo 2. Using the short declaration Operator (:=)
?  age := 21
*/

func main() {
	var x int = 7
	var y = "manish"
	lastname := "kumar"

	car, cost := "BMW", 570000 // multipal variable declaration
	fmt.Println(car, cost)

	var opend = false
	opend, file := true, "new.txt"
	_, _ = opend, file
	fmt.Println("this is variable", x, y, lastname)

	// multipal declaration

	var (
		salary    float64
		firstname string
		gender    bool
	)

	fmt.Println(salary, firstname, gender)

	var a, b, c int
	fmt.Println(a, b, c)
	// multipal assignment

	var i, j int
	i, j = 10, 20
	fmt.Println(i, j)

	//swap two number

	j, i = i, j
	fmt.Println(i, j)

	sum := 5 + 3.5
	fmt.Println(sum)

	// you can not assign float value in int

	var one = 4
	var two = 7.5

	// one = two  you cant not siign float value in int
	// one = 7.5 // error: cannot use b (type float64) as
	one = int(two)
	fmt.Println(one, two)

	// ! GO Zero Values
	// ?GO has zero values for all types. Zero values are the default values assigned to variables when
	// they are declared without an explicit initial value.
	// Zero values are as follows:
	// !1. Boolean: false
	// !2. Numeric: 0
	// !3. string ""(empty string)
	// !4. Complex: 0+0i
	// !5. pointer type : nil

	// ? Go is a statically and strong typed programming language

	var aa int = 10
	var bb = 15.5
	cc := "Go Pro"

	_, _, _ = aa, bb, cc

	fmt.Println(aa, bb, cc)

	/*
		! -------------NAMING CONVENTION---------------------------

		   todo =  Naming (variable, function, packages) Conventions is GO
		   ?  nameing Convention are important for code readability and maintainability.
		   * Names start with a letter or an underscore(_)
		   * Case Matter: quickSort and QuickSort are diffrent variable
		   * Go keywords (25) can not be used as names
		   * Use the first letters of the words


		   !2. Variable Names
			* Use camelCase for variables and function names.
			* Variable names should be concise but descriptive enough to understand the intent. Avoid one-letter names unless in small scopes like loop indices.
		*var userName string  // Good
		! var uName string     // Avoid (too unclear)
		* var x int            // Acceptable in small contexts like loops

		! 1. Package Names
			*Package names should be short, lowercase, and should not contain underscores or mixed case.
		?package math  // Good
		?package http  // Good
		!package utilities  // Avoid
		!package fileHelper  // Avoid
	*/


}
