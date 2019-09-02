package main
// If and Switch statements
// If statements: Operators, If-else and if-else if statements
// Switch statements: Simple cases, Cases with multiple tests, Falling through, Type switches
import (
	"fmt"
)
func main() {
	if true {
		fmt.Println("The test is true")
	}

	// Initializer syntax
	if equal := 3 == 3; equal {
		fmt.Println("Two numbers are equal")
	}

	// Guess game
	number := 30
	guess := 20
	if guess == number {
		fmt.Println("You got it!")
	}
	if guess < number {
		fmt.Println("Too low!")
	}
	if guess > number {
		fmt.Println("Too high!")
	}
	if guess < number || guess > number {
		fmt.Println("Your guess is wrong!")
	}

	// if - else
	if guess == number {
		fmt.Println("The guess is correct!")
	} else {
		fmt.Println("The guess is wrong!")
	}

	// switch statements
	// switch statement does not fall through in Go
	// cases have to be distinct
	switch i := 5; i {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("another number")
	}

	// Tag-less syntax
	// break is implicitly applied
	i := 10
	switch {
	case i <= 10:
		fmt.Println("less than or equal to ten")
		fallthrough // directly get into the next case, no matter the logical test
	case i >= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}

	// type switch
	var t interface{} = 1
	switch t.(type) {
	case int:
		fmt.Println("t is an int")
		break // this can be used to exit the switch statement
		fmt.Println("this is a test")
	case string:
		fmt.Println("t is a string")
	default:
		fmt.Println("t is another type")
	}
}