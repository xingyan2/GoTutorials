package main

import (
	"fmt"
)
/**
Basic syntax, Parameters, Return values, Anonymous functions, Functions as types, Methods
 */
func main() {
	sayMessage("This is your number:", 10)
	sayGreeting("Good Morning", "Stephen")
	firstName := "Peter"
	lastName := "Parker"
	sayNames(&firstName, &lastName)  // pointers as arguments

	sum(1,2,3,4)
	product := product(1,2,3,4,5)
	fmt.Println("The product is:", *product)
	res := newSum(1,2,3,4,5)
	fmt.Println("The new sum is:", res)

	quo, err := divide(5.0, 3.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("The result is:", quo)

	// More about functions
	// Anonymous function
	func() {
		fmt.Println("This is an anonymous function")
	}()  // the following parentheses invoke the function immediately
	// The anonymous function can be helpful as follows
	// For asynchronous execution
	for i := 0; i < 3; i++ {
		func(i int) {
			fmt.Println(i)
		}(i)
	}

	// Function as variable
	var f func(float64, float64) (float64, error)
	f = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("cannot divide by zero")
		} else {
			return a / b, nil
		}
	}
	d, err2 := f(5.0, 3.0)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println(d)

	// Methods
	g := greeter{
		greeting: "Hello",
		name: "Go",
	}
	g.greet()
}

func sayMessage(msg string, index int) { // the function takes two parameters of different types
	fmt.Println(msg, index)
}

func sayGreeting(greeting, name string) { // type can be declared once for multiple parameters
	fmt.Println(greeting, name)
}

func sayNames(firstName, lastName *string) {  // Pointers are passed as arguments, can be more efficient.
	fmt.Printf("The character's name is %v %v\n", *firstName, *lastName)
	*firstName = "Tony"
	*lastName = "Stark"
	fmt.Printf("The new character's name is %v %v\n", *firstName, *lastName)
}

func sum(values ...int) int {  // this variadic function takes any number ints as arguments, and it returns an int
	var res int = 0
	for _, v := range values {
		res += v
	}
	fmt.Println("The sum is:", res)
	return res  // the value to be returned
}

func product(values ...int) *int {  // this function returns a pointer pointing to int
	var res int = 0
	for _, v := range values {
		res += v
	}
	return &res
	// This pointer is supposed to be associated with memory space of the function's execution stack
	// when the function exits, the variable gets promoted to the heap.
}

func newSum(values ...int) (res int) {  // Named return values (syntactic sugar for declaring variables)
	for _, v := range values {
		res += v
	}
	return
}

func divide(a, b float64) (float64, error) {  // You can return multiple values
	if b == 0.0 {
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

type greeter struct {
	greeting string
	name string
}
// definition of method
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.name = ""  // this is a value receiver, just a copy.
	// but this will be different if we write it as: func(g *greeter) greet(){}
}