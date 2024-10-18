**Init**

In Go, the init function is a special function that is executed automatically when a package is initialized, before the main function is called. It is primarily used to set up any necessary initialization logic for the package, such as setting up variables or executing code that should run prior to the program’s main execution.<br>

Automatic Execution: The init function is called automatically by the Go runtime; you don't need to call it explicitly.<br>

Package Initialization: Each package can have its own init function. If a package imports other packages, the init functions in those imported packages are called before the init function of the importing package.<br>

No Arguments and No Return Values: The init function does not take any parameters and does not return any values. You can have multiple init functions within a single package, and they will be executed in the order they appear.<br>

Precedes main: The init function is executed before the main function of the program, ensuring that any setup needed for the program's execution is completed first.<br>

***Example of init in Go:***

package main

import (
    "fmt"
)

// This is the init function for the package
func init() {
    fmt.Println("Initializing the package...")
}

// The main function of the program
func main() {
    fmt.Println("Main function called.")
}

