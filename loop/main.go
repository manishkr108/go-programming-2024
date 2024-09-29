package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)

	}

	//! Go dont have while loop below code is same effect as while loop
	j := 10
	for j > 10 {
		fmt.Println(j)
		j--
	}
	// Infinie loop
	// sum := 0

	// for {
	// 	sum++
	// }
	// fmt.Println(sum)

	// * Continue Statement

	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}

	// ! Breck Statement

	count := 0

	for i := 0; true; i++ {
		if i%13 == 0 {
			fmt.Printf("%d divided by 13\n", i)
			count++
		}
		if count == 10 {
			break // breack statement is not teminate program intirely it is jumping out of the loop and program continue at the statement following curlybreckets

		}
	}

	fmt.Println("just a message")

	// !--------LEBEL Statement ---------------

	people := [5]string{"hey", "mani", "kumar", "bangalore", "all good"}
	frends := [2]string{"mani", "raj"}

outer:
	for index, name := range people {
		for _, frnd := range frends {
			if name == frnd {
				fmt.Printf("a friend %q at index %d\n", frnd, index)
				break outer
			}
		}
	}
	fmt.Println(" Next instraction after loop")

	//! GOTO

	ii := 0
loop:
	if ii < 5 {
		fmt.Println(ii)
		ii++
		goto loop
	}

	//! Switch Statement

	language := "golang"
	switch language {
	case "java":
		fmt.Println("java is a language")
	case "c":
		fmt.Println("c is a language")
	case "golang":
		fmt.Println("golang is a language")
	default:
		fmt.Println("language not found")
	}

	n := 5
	switch true {
	case n%2 == 0:
		fmt.Println("Even Number")
	case n%2 != 0:
		fmt.Println("Odd Number")
	}
}
