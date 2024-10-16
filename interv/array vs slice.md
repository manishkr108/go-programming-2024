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

<hr>
Slice
- Has a dynamic length (it can  shink or grow)
- The length of a slice is not part of its type, it is determined at runtime
- an uninitialized slice is equal to nil (its zero value is nil)
--------------------------------------------------------------<br>
<li>both slice and an array can contain only same type of elements</li>
<li> we can create a keyed slice like keyed array</li>

<b>In Go, list := []int{1, 2, 3, 4} is a slice, not an array.<b>

Here's the difference between a slice and an array in Go:

Array: An array has a fixed size, which is part of its type, and it cannot be resized. For example:<br>


var arr [4]int = [4]int{1, 2, 3, 4}  // This is an array with a fixed length of 4 <br>
Slice: A slice is a dynamically sized, flexible view into the elements of an array. It does not include the size in its type, and you can append or modify it as needed. For example:<br>


list := []int{1, 2, 3, 4}  // This is a slice
Since your declaration omits the length ([]int{}), list is a slice.

In the declaration list := [...]int{2, 3, 4, 5}, this is an array, not a slice.

The [...] syntax in Go is used to define an array where the length is inferred from the number of elements provided in the literal. In this case, the array has a fixed size of 4 elements.

Explanation:
Array: The type [...]int{2, 3, 4, 5} means the array has a length of 4, and it's determined by the number of elements between the curly braces. This is a fixed-length array.


list := [...]int{2, 3, 4, 5}  // Array of length 4
Slice: A slice would not specify the length or use [...]. For example:


list := []int{2, 3, 4, 5}  // This would be a slice
Since your code uses [...], list is an array.



Here's the comparison between arrays and slices in Go in a table format:

| **Feature**               | **Array**                                                                 | **Slice**                                                             |
|---------------------------|---------------------------------------------------------------------------|----------------------------------------------------------------------|
| **Size**                  | Fixed size, determined at compile time.                                    | Dynamic size, can grow or shrink.                                     |
| **Syntax**                | `var arr [4]int = [4]int{1, 2, 3, 4}`                                      | `var slice []int = []int{1, 2, 3, 4}`                                 |
| **Length**                | Length is part of the type (e.g., `[4]int`).                               | Length is not part of the type (e.g., `[]int`).                       |
| **Capacity**              | Capacity is always equal to length.                                        | Capacity can be larger than the current length.                       |
| **Memory Allocation**     | Allocates memory for all elements, even if unused.                         | Uses dynamic arrays under the hood, reallocating as needed.           |
| **Resizable**             | Cannot change size once created.                                           | Can be resized using functions like `append()`.                       |
| **Reference Behavior**    | Direct value type (copy of the array).                                     | Reference type, referencing the underlying array.                     |
| **Efficiency**            | More memory-efficient if the fixed size is known.                          | More flexible and convenient for variable-sized data.                 |
| **Passing to Functions**  | Passed by value (copy of the array).                                       | Passed by reference (points to the underlying array).                 |
| **Built-in Functions**    | No built-in functions like `append`, `copy`.                               | Has built-in functions such as `append`, `copy`, `len`, `cap`.        |


Key Features of Slices:
Dynamic size: The size of a slice can change as elements are added or removed.

Underlying array: A slice points to a portion of an underlying array.

Capacity and length: Slices have both length (the number of elements in the slice) and capacity (the total number of elements the underlying array can hold).

package main

import "fmt"

func main() {
    // Create a slice with 3 elements
    nums := []int{10, 20, 30}  // Slice literal (dynamic array)
    
    fmt.Println("Initial slice:", nums)   // Output: [10, 20, 30]

    // Append new elements to the slice
    nums = append(nums, 40, 50)
    fmt.Println("After appending:", nums) // Output: [10, 20, 30, 40, 50]

    // Slice an existing slice (create a sub-slice)
    subSlice := nums[1:4]   // Slicing from index 1 to 3 (not including 4)
    fmt.Println("Sub-slice:", subSlice)   // Output: [20, 30, 40]

    // Modify an element of the sub-slice (this affects the original slice)
    subSlice[0] = 25
    fmt.Println("Modified sub-slice:", subSlice) // Output: [25, 30, 40]
    fmt.Println("Original slice after modification:", nums) // Output: [10, 25, 30, 40, 50]

    // Use the built-in `len` and `cap` functions
    fmt.Println("Length of slice:", len(nums))   // Output: 5
    fmt.Println("Capacity of slice:", cap(nums)) // Output: 6 (capacity can be larger than the length)
}

Explanation:
Slice Literal: nums := []int{10, 20, 30} creates a slice with 3 elements. The underlying array is automatically managed by Go.<br>
Appending: append(nums, 40, 50) adds two new elements to the slice, dynamically increasing its size.<br>
Slicing: subSlice := nums[1:4] creates a sub-slice from the second to the fourth element of nums. The sub-slice is a reference to the same underlying array, so changes in the sub-slice affect the original slice.<br>
Modifying a Sub-slice: Changing subSlice[0] updates the original nums slice, showing the reference behavior of slices.<br>
Length and Capacity: len(nums) gives the current number of elements, and cap(nums) gives the capacity (which can be larger than the length due to underlying array resizing).<br>

<strong>Key Points About Slices:</strong>
Slicing: You can create a new slice from an existing slice using the slicing operator slice[low:high].<br>
Capacity: Slices may have extra capacity. As you append elements, Go automatically resizes the underlying array when necessary.<br>
Zero Value: The zero value of a slice is nil, which behaves like a slice with length and capacity 0.<br>


<b>Recap of Key Functions:<b>
append(): Adds elements to a slice, dynamically growing it.<br>
copy(): Copies elements from one slice to another.<br>
len(): Returns the current number of elements in the slice.<br>
cap(): Returns the total capacity of the slice, or how many elements it can hold before resizing.<br>

Here’s a simple yet powerful example to supercharge your Go skills with slices👇:

tasks := []string{"Design", "Develop", "Test"}
tasks = append(tasks, "Deploy")
backupTasks := make([]string, len(tasks))
copy(backupTasks, tasks)
fmt.Println(tasks, backupTasks)

