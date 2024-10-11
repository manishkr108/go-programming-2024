package main

import "fmt"

func main() {
	var cities []string
	fmt.Println("Cities is equal to nil", cities == nil) //Cities is equal to nil true
	fmt.Printf("cities %#v\n", cities)                   //cities []string(nil)
	fmt.Println(len(cities))                             //0

	//cities[0] = "ABC" // this is error panic: runtime error: index out of range [0] with length 0

	// keep in mind that the elements have a position or index that starts from zero not one

	numbers := []int{2, 3, 4, 5} // right side of eqil = is the slice literal
	fmt.Println(numbers)         // [2 3 4 5]
	// another way to clreate slice is to use the make function

	num := make([]int, 2)
	fmt.Printf("%#v\n", num) // []int{0, 0}

	type names []string

	friends := names{"manish", "raj"}
	fmt.Println("my best friends is ", friends) //y best friends is  [manish raj]
	friends[0] = "priya"
	fmt.Println("my best friends is ", friends) //y best friends is  [priya raj]

	// number slice
	for index, value := range numbers {
		fmt.Printf("index %d value %d\n", index, value)
	}

	var n []int = numbers
	// n = numbers
	// `n = numbers` is assigning the slice `numbers` to the slice `n`. This means that both `n` and
	// `numbers` will refer to the same underlying array. Any changes made to the elements of the slice
	// through one variable will be reflected in the other variable as well since they are pointing to the
	// same data.
	fmt.Println(n)

	var mn []int
	fmt.Println(mn == nil) //true

	mk := []int{}
	fmt.Println(mk == nil) //false

	a, b := []int{1, 2, 3}, []int{1, 2, 3}
	//fmt.Println(a == b) // this is error // slice only compire nil

	var eq bool = true

	a = nil

	if len(a) != len(b) {
		eq = false
	}

	for i, valueA := range a {
		if valueA != b[i] {

			eq = false
			break
		}

	}

	if eq {
		fmt.Println("a and b slice are equal")
	} else {
		fmt.Println("a and b slice are not equal")
	}

	// How to append slice
	// append does't modify the initial slice it return a brand the new one;
	nm := []int{2, 3}
	nm = append(nm, 4)
	fmt.Println(nm) // [2 3 4]
	// other example

	nms := []int{3, 4, 5, 6}
	nms = append(nms, 10, 11, 12)
	fmt.Println(nms)

	nn := []int{10, 20}
	nm = append(nm, nn...)
	fmt.Println(nm) //[2 3 4 10 20]
	// The append() function returns-iohg

	//Copy function
	// the copy function does't always clone or dublicate a slice
	src := []int{10, 20, 30, 40}
	dist := make([]int, len(src))
	q := copy(dist, src)
	fmt.Println(src, dist, q)

	/*
			Explanation:
		src := []int{10, 20, 30, 40}: The src slice is initialized with 4 elements: [10, 20, 30, 40].

		dist := make([]int, len(src)): The dist slice is created with a length equal to src, but all elements are initialized to 0. So initially, dist = [0, 0, 0, 0].

		q := copy(dist, src): The copy() function copies elements from the src slice into the dist slice. Since both slices have the same length, all 4 elements will be copied. The copy() function returns the number of elements copied, which in this case is 4.

		fmt.Println(src, dist, q): This prints:

		The original src slice.
		The dist slice, which now contains the copied elements from src.
		The value of q, which represents the number of elements copied.

		/ description/
		src := []int{10, 20, 30, 40}        // Initialize a slice 'src' with 4 elements
		dist := make([]int, len(src))       // Create a 'dist' slice with the same length as 'src', but all elements initialized to 0
		q := copy(dist, src)                // Copy elements from 'src' to 'dist'. 'q' holds the number of elements copied
		fmt.Println(src, dist, q)           // Print 'src', 'dist', and the number of elements copied ('q')
	*/

	//Slice Expressions
	// slice does't modify the array or slice, it return a new one.
	ab := [5]int{1, 2, 3, 4, 5}
	// aa[start:stop]
	bb := ab[1:3]
	fmt.Printf("%v, %T\n", bb, bb) // [2 3] []int
	// we notice that slicing an array returns a slice , not array
	q1 := []int{1, 2, 3, 4, 5, 6}
	q2 := q1[1:3]
	fmt.Printf("%v, %T\n", q2, q2) // [2 3]
	fmt.Println(q1[2:])            //[3 4 5 6] same as q1[2:len(q1)]
	fmt.Println(q1[:3])            //[1 2 3] same as q1[0:3]
	fmt.Println(q1[:])             //[1 2 3 4 5 6] same as q1[0:len(s1)]
	// fmt.Println(q1[:100]) // this is error index out of bound
	q1 = append(q1[:4], 100)
	fmt.Println(q1) //[1 2 3 4 100]

}
