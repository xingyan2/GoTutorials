package main

import (
	"fmt"
)
// Creating pointers, Dereferencing pointers, The new function, Working with nil, Types with internal pointers
func main() {
	var a int = 42
	var b *int = &a
	fmt.Println(&a, b)
	fmt.Printf("The value pointed to by b %v\n", *b)

	// Pointer arithmetic is not explicitly supported by Go
	a1 := [3]int{1,2,3}
	b1 := &a1[0]
	c1 := &a1[1]
	fmt.Printf("%v %p %p\n", a1, b1, c1)

	var ms *myStruct  // a pointer that points to myStruct data type
	fmt.Println(ms)  // initialized to nil by default
	ms = new(myStruct)
	ms.foo = 42  // Compiler understands this as (*ms).foo = 42
	fmt.Println(ms.foo)

	// Be careful when operating on slices and maps, especially when you pass them around and modify their values
}

type myStruct struct {
	foo int
}
