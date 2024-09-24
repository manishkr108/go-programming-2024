package main

import "fmt"

/*
! ------------------Go Data Types-1------------
* A type determine a set of value together with operation and methods specific to those value
* There are //! predeclear type, introduced type with type declarations and composite types
* Types: array, slice, map, struct , pointer, function , interface and channel types

!------Predeclared, Build-in Type---------------
* Numeric Type
? + int8, int16, int32, int64, uint8, uint16,
? uint32, uint64, uintptr, int, uint, float32, float64, complex64
? complex128, byte, rune, string, bool, error

*/
func main() {
	var i1 int8 = 100 //  ! the minimum value of int8 is -128
	fmt.Printf("%T\n", i1)

	var i2 uint16 = 65535
	fmt.Printf("%T\n", i2) //uint16

	//* Float Type
	var f1, f2, f3 float64 = 1.1, -0.2, 5.
	fmt.Printf("%T %T %T\n", f1, f2, f3) //float64 float64 float64

	//* Golang does't have char data type , it uses byte and rune represent char values
	//!RUNE Type

	var r rune = 'f'
	fmt.Printf("%T\n", r) //int32
	fmt.Println(r)        //102
	fmt.Printf("%x\n", r) //66

	//! String Type
	var s string = "Hello, Go"
	fmt.Printf("%T\n", s) //string

	//* Array Type
	var arr [5]int = [5]int{1, 2, 3, 4}
	fmt.Printf("%T\n", arr)

	// ! Slice Tyep
	var city = []string{"london", "Tokyo", "New York", "4"}
	fmt.Printf("%T\n", city) //[]string

	// * Map Tyep
	frout := map[string]float64{
		"apple":  1.2,
		"banana": 2.3,
		"orange": 3.4,
	}
	fmt.Printf("%T\n", frout) //map[string]float64

	// ! Struct Type
	type person struct {
		name string
		age  int
	}
	var you person
	fmt.Printf("%T\n", you) //main.person

	// ! Pointer Type
	var xx int = 2
	ptr := &xx
	fmt.Printf("ptr is of type %T with the value of %v\n", ptr, ptr) //*int

	var golang string = "😊😊"
	fmt.Printf("%v\n", golang) //string

}
