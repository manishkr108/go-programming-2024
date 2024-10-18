<h1>What is struct</h1>

-> struct (short for "structure") is a composite data type that groups together different variables under a single name.<br>
->Each of these variables can be of different types, and they are called fields of the struct.<br>
-> In Go, Struct are one of the most powerful and commonly used data type<br>
-> Struct allow you to group related data together, making your code more orgnized and easy to work with.<br>

type Person struct {
    Name    string
    Age     int
    Address string
}

main
person1 := Person{
        Name:    "John Doe",
        Age:     30,
        Address: "123 Main Street",
    }

 // Accessing the fields
    fmt.Println("Name:", person1.Name)
    fmt.Println("Age:", person1.Age)
    fmt.Println("Address:", person1.Address)

    // Modifying a field
    person1.Age = 31
    fmt.Println("Updated Age:", person1.Age)

A struct is a callection of fields, where each field has its own type.<br>
it's similar to classes in other programming language, but without method attached.

<h3>Syntex</h3>

type Person struct {
    Name string
    Age int
}


type Event struct {
	ID          int64     // id of the event
	Name        string    `json:"name" binding:"required"` // name of the event
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`  // location of the event
	StartTime   time.Time `json:"startTime" binding:"required"` // start time of the event
	UserID      int64     `json:"user_id"`
}

<h3>Anonymous Structs</h3>
- Go also allow the use of anonymous structs, which does't require pre-defined type<br>
- You can define and use a struct without explicitly creating a named type. This is useful for quick one-off uses where you don’t need to reuse the struct elsewhere.<br>

**example**
person := struct {
    Name string
    Age  int
}{
    Name: "Alice",
    Age:  25,
}

fmt.Println(person.Name) // Output: Alice

**3. Embedded Struct (Struct Composition)**

Go doesn’t support inheritance, but you can achieve a similar effect using struct embedding. This allows one struct to contain another, effectively reusing fields from the embedded struct.<br>

**Example**
type Address struct {
    City  string
    State string
}

type Person struct {
    Name    string
    Age     int
    Address // Embedded struct
}

func main() {
    person := Person{
        Name: "John",
        Age:  30,
        Address: Address{
            City:  "New York",
            State: "NY",
        },
    }

    fmt.Println(person.City)  // Output: New York
    fmt.Println(person.State) // Output: NY
}


**4. Struct with Methods**
-Go allows you to associate methods with a struct, enabling you to create behavior associated with the data within the struct.<br>

type Person struct {
    Name string
    Age  int
}

// Method associated with the Person struct
func (p Person) Greet() {
    fmt.Printf("Hello, my name is %s.\n", p.Name)
}

func main() {
    person := Person{Name: "Bob", Age: 25}
    person.Greet() // Output: Hello, my name is Bob.
}

**5. Pointer to a Struct**
-Go allows you to use pointers with structs. This is particularly useful for modifying the struct's data or avoiding the cost of copying large structs.<br>

type Person struct {
    Name string
    Age  int
}

func modifyPerson(p *Person) {
    p.Age = 30
}

func main() {
    person := Person{Name: "Alice", Age: 25}
    modifyPerson(&person)
    fmt.Println(person.Age) // Output: 30
}

**6. Tagged Structs**
You can add tags to struct fields, often used in conjunction with serialization (like json, xml, etc.). These tags help control how data is marshaled or unmarshaled.<br>

package main

import (
    "encoding/json"
    "fmt"
)

// User struct with a JSON tag for username
type User struct {
    Username string `json:"username"` // JSON tag
}

func main() {
    // Create an instance of User
    user := User{
        Username: "john_doe",
    }

    // Marshal the struct to JSON
    jsonData, err := json.Marshal(user)
    if err != nil {
        fmt.Println("Error marshaling to JSON:", err)
        return
    }

    // Print the JSON output
    fmt.Println(string(jsonData)) // Output: {"username":"john_doe"}
}



