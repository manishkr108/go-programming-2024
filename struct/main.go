package main

import (
	"encoding/json"
	"fmt"
)

// 1. Defining a struct
type Person struct {
	Name string `json:"name"` // 6. Struct tags (for JSON encoding)
	Age  int    `json:"age"`
}

// 5. Embedding in Struct (Composition)
type Employee struct {
	Person   // Embedding Person struct
	Position string
	Salary   float64
}

// 4. Pointer to Struct (function to update age)
func UpdateAge(p *Person, newAge int) {
	p.Age = newAge
}

func main() {
	// 2. Initializing a Struct (different ways)

	// Using field names
	person1 := Person{Name: "John", Age: 30}

	// Without field names (order matters)
	person2 := Person{"Alice", 25}

	// Zero value initialization
	var person3 Person // Fields will have default values (Name: "", Age: 0)

	// 3. Accessing and Modifying Struct Fields
	fmt.Println("Initial Name:", person1.Name) // Accessing field
	person1.Age = 31                           // Modifying field
	fmt.Println("Updated Age:", person1.Age)

	// 4. Passing Struct by Reference
	fmt.Println("Before UpdateAge:", person2.Age)
	UpdateAge(&person2, 26) // Passing a pointer to the struct
	fmt.Println("After UpdateAge:", person2.Age)

	// 5. Embedding a Struct (Composition)
	employee := Employee{
		Person:   Person{Name: "Bob", Age: 40},
		Position: "Manager",
		Salary:   75000.50,
	}
	fmt.Printf("Employee: %s, Position: %s, Salary: %.2f\n", employee.Name, employee.Position, employee.Salary)

	// 6. Struct Tags (JSON Encoding Example)
	jsonData, _ := json.Marshal(person1) // Marshal person1 to JSON
	fmt.Println("JSON Output:", string(jsonData))

	// 7. Zero Value Struct
	fmt.Printf("Zero Value Struct: %+v\n", person3) // Print zero-initialized struct

	// 8. Comparing Structs
	person4 := Person{Name: "John", Age: 31}
	if person1 == person4 {
		fmt.Println("person1 and person4 are equal!")
	} else {
		fmt.Println("person1 and person4 are not equal.")
	}
}
