package main

import (
	"fmt"
)

// Defer
// Panic
// Recover

func main() {
	// Defer
	fmt.Println("start")
	//defer fmt.Println("middle")
	//defer fmt.Println("end")  // multiple defers follow LIFO. "start" "end" "middle"

	// defer keyword can be used to associate the opening of a resource with its closing
	// but defer statements don't execute until functions exit
	// this may cause problems if too many resources were opened and are to be closed at the end

	//a := "start"
	//defer fmt.Printf("The value of a is: %v\n", a)  // it takes the value when defer is called
	//a = "end"

	// Panic
	fmt.Println("start")
	defer fmt.Println("this was deferred")
	panic("something bad happened")  // panic comes after defer
	fmt.Println("end")
}
