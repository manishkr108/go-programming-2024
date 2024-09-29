package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
! how to get user input from command line
!  go provide a package like os
!  os.Stdin is the standard input stream

*/

func main() {

	// fmt.Println("os.Args", os.Args)
	// fmt.Println("Path : ", os.Args[0])
	// fmt.Println("First Arg : ", os.Args[1])
	// fmt.Println("Second Arg : ", os.Args[2])
	// fmt.Println("Length of Args : ", len(os.Args))

	var result, err = strconv.ParseFloat(os.Args[1], 64)

	fmt.Printf("%T\n", os.Args[1])
	fmt.Printf("%T\n", result)
	_ = err
	//! SIMPLE STATEMENT

	i, err := strconv.Atoi("45")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(i)

	if i, err := strconv.Atoi("20"); err == nil {
		fmt.Println("no error, i is", i)
	} else {
		fmt.Println(err)
	}
}
