<h2>Golang Array vs Slice</h2>

<h3>1. Array in Go</h3>
<strong>Fixed Size:</strong> Arrays in Go have a fixed size defined at the time of declaration, and this size cannot change.<br>
<strong>Value Type:</strong> Arrays are value types, meaning that when you pass an array to a function, it gets copied.<br>
<strong>Memory:</strong> Arrays allocate memory for all elements at once, even if they are not used.<br>


----------Array in Go --------------

1. An array is a composit, indexable type that stores a collection of elements of same type.
2. An array has "fixed length"
3. Every element in an array must be of same type
4. Go store the elements of the array in contiguous memory locations and this way its very efficient.
5. The length and the elements type determines the type of an array. "the length belongs to array type" and its determined at compile time

example 
accounts := [3]int{50,60,70}

the array called accounts that consists of 3 integers has it type [3]int.

func main() {
	var numbers [5]int
	fmt.Printf("%v\n", numbers)//[0 0 0 0 0]
	fmt.Printf("%#v\n", numbers)//[5]int{0, 0, 0, 0, 0}
}

=> range = here range is just a language key word used for iteration.


// 2 arrays are equal if they have the same length and the same elements, in the same order

[1,2,3] == [1,2,3] = true
[1,2,3]==[1,3,2] = false




