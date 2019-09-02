package main
/*
Boolean type
Numeric type:
	Integers, Floating point, Complex numbers
Text type
*/

import (
	"fmt"
)

func main() {
	// Booleans
	var n bool = true  // Declare a Boolean variable
	fmt.Printf("%v, %T\n", n, n)

	m := 1 == 1  // Logical test
	fmt.Printf("%v, %T\n", m, m)

	var p bool // Uninitialized boolean variable is false by default
	fmt.Printf("%v, %T\n", p, p)

	// Integers
	// int8, int16, int32, int64
	var a int = 15
	fmt.Printf("%v, %T\n", a, a)
	var b uint16 = 40
	fmt.Printf("%v, %T\n",b, b)
	c := 4
	fmt.Println(a * c * int(b))  // type conversion before doing operations

	// Bit operators
	d := 10 // 1010
	e := 3  // 0011
	fmt.Println(d & e)  // 0010  AND
	fmt.Println(d | e)  // 1011  OR
	fmt.Println(d ^ e)  // 1001  XOR
	fmt.Println(d &^ e) // 0100  AND NOT: NOT first, then AND
	fmt.Println(e << 2) // left shift  1100
	fmt.Println(e >> 1) // right shift 0001

	// Floating point numbers
	f := 3.14
	fmt.Printf("%v, %T\n", f, f)
	f = 13.7e32
	fmt.Printf("%v, %T\n", f, f)
	f = 2.1E14
	fmt.Printf("%v, %T\n", f, f)

	// Complex numbers
	var g complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", g, g)
	var h complex64 = 2 + 1i
	fmt.Printf("%v, %T\n", real(h), real(h))  // get the real part of complex number h
	fmt.Printf("%v, %T\n", imag(h), imag(h))  // get the imaginary part of complex number h
	// complex number operations
	fmt.Println(g * h)
	fmt.Println(g + h)
	// Create a new complex number
	var i complex64 = complex(5, 12)
	fmt.Printf("%v, %T\n", i, i)  // get the imaginary part of complex number h

	// Text types
	// string is a sequence of bytes
	s := "this is a string"
	s2 := "another string"
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v, %T\n", s + s2, s + s2)  // string concatenation

	by := []byte(s)  // convert string to byte slice
	fmt.Printf("%v, %T\n", by, by)

	// rune type is an alias for int32, and is meant to represent a Unicode CodePoint
	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r)
}
