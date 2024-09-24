package main

import (
	"fmt"
)

/*
! IOTA
* within a constant decleration, the predecleared identifier iota represents successive untyped integer constants.

*/

func main() {
	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)

	fmt.Println(c1, c2, c3) // 0 1 2

	const (
		North = iota
		South
		East
		West
	)

	fmt.Println(North, South, East, West) //0 1 2 3

	const (
		a = (iota * 2) + 1
		b
		c
	)

	fmt.Println(a, b, c) //1 3 5

	const (
		i = iota + 5
		j
		k
	)

	fmt.Println(i, j, k) //5 6 7

}
