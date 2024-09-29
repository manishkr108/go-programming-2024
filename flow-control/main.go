package main

import "fmt"

/*
 ! flow control if , else if, else statement
 * if, else if, and else statement that are used for decision making



*/

func main() {

	// if statement
	price, inStock := 100, true
	if price > 80 {
		fmt.Println("Price is greater than 80")
	}
	_ = inStock

	if price <= 100 && inStock {
		fmt.Println("buy it!")
	}
	// this is error in golang
	// if price{
	// 	fmt.Println("price is true")
	// }

	if price <= 100 {
		fmt.Println("price is less than or equal to 100")
	} else if price == 80 {
		fmt.Println("On the edge")
	} else {
		fmt.Println("price is greater than 100")
	}

	age := 50

	if age >= 0 && age < 18 {
		fmt.Printf("you can not vote! please return in %d years !\n", 18-age)
	} else if age == 18 {
		fmt.Println("you can vote , this is your first vote")
	} else if age > 18 && age <= 100 {
		fmt.Printf("you can vote , its important\n")
	} else {
		fmt.Println("invalid age")
	}
}
