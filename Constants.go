package main
import (
	"fmt"
)
/*
Naming convention
Typed constants
Untyped constants
Enumerated constants
Enumeration expressions
*/

const a int16 = 27

// a constant block
const (
	a1 = iota  // iota keyword represents successive integer constants 0, 1, 2, ...
	a2 = iota  // it resets to 0 whenever the word const appears in the source code
	a3 = iota  // and increments after each const specification
)

const (
	a4 = iota
	a5		// compiler can infer the pattern of assignments
)

const (
	errorSepcialist = iota
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

const (
	_ = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
)

const (
	isAdmin = 1 << iota
	isHeadquaters
	canSeeFinancials
)

func main() {
	// lowercase first letter indicates that it will not be exported
	const myConst int = 42
	// myConst = 32  you cannot change the value of a constant
	// const myConst2 float64 = math.Sin(1.58)  You cannot set the constant equal
												// to a value that needs to be determined at runtime
	fmt.Printf("%v, %T\n", myConst, myConst)

	const a int = 14
	fmt.Printf("%v, %T\n", a, a) // shadowing the constant a

	const b = 42
	fmt.Printf("%v, %T\n", b, b) // the compiler can infer the type of the constant
	var c int16 = 21
	fmt.Printf("%v, %T\n", b + c, b + c) // compiler infers that b is of type int16 in this case

	// Enumerated constants
	fmt.Printf("%v\n", a1)
	fmt.Printf("%v\n", a2)
	fmt.Printf("%v\n", a3)

	fmt.Printf("%v\n", a4)  // from another const block
	fmt.Printf("%v\n", a5)

	// application of enumerated constants
	var specialistType int  // specialistType is initialized to 0 by default
	fmt.Printf("%v\n", specialistType == catSpecialist)

	fileSize := 4000000000.  // floating point by adding a dot .
	fmt.Printf("%.2fGB\n", fileSize/GB)

	var role byte = isAdmin | canSeeFinancials
	fmt.Printf("%b\n", role)
	fmt.Printf("Is Admin? %v\n", isAdmin & role == isAdmin)
}
