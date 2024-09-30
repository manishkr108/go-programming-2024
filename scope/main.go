package main

/*
! Scope means visibility.
	* the scope or the lifetime of a variable is the interval of time during wh
*/

import (
	"fmt"
	// import "fmt" // this is error

	f "fmt"
)

// it permitted in go

const done = false // package scope

var b int = 10 // if i am not use this variable there is no error
// it can be used also in other files of the same package

func main() {

	var task = "running" // local (block) scope

	fmt.Println(task, done)
	const done = true                        // local scope
	f.Printf("done in main() is %v\n", done) // it permitted

	fun() // done in  fun(): false
	f.Println("i am using the fmt alias hear")

}

func fun() {
	fmt.Printf("done in  fun(): %v\n", done) // this done from package scope
	a := 10                                  // if i am not use a varibale it will return error unused variable
	_ = a                                    // the blank identifire is represent by the underscore charactor and can be used to mute the compile time error returned

}
