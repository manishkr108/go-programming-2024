package main

import "fmt"

func main() {
	name, age := "manish", 30
	fmt.Printf("%s is %d \n", name, age)

	fmt.Printf("%[2]s is %[1]d\n", age, name)

	fmt.Printf("%[1]T (%[1]v)\n", name)
	fmt.Printf("%[1]T (%[1]v)\n", age)

	fmt.Printf("My name is %s. Yes, you heard that right: %s\n", name, name)

	fmt.Printf("My name is %[1]s. Yes, you heard that right: %[1]s\n", name)

	number := 10000000
	better := 10_000_000

	fmt.Println(number == better)
}
