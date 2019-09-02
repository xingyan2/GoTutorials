package main
import (
	"fmt"
)
// Arrays: Creation, Built-in functions, Working with arrays
// Slices: Creation, Built-in functions, Working with slices

func main() {
	grades := [...]int{97, 85, 93}  // create and initialize a new array
	var students [3]string
	students[0] = "Emma"
	students[1] = "Peter"
	students[2] = "Jeremy"
	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Number of students: %v\n", len(students)) // get the length of an array

	// 2D Array
	var identityMatrix [3][3]int = [3][3]int{[3]int{1,0,0}, [3]int{0,1,0}, [3]int{0,0,1}}
	fmt.Println(identityMatrix)
	identityMatrix[0] = [3]int{1,0,0}

	// How values are passed around
	a := [...]int{1,2,3}
	b := a  // b is assigned a copy of a
	b[1] = 5
	fmt.Println(a)  // [1 2 3]
	fmt.Println(b)  // [1 5 3]

	// Slices
	a1 := []int{1,2,3}
	fmt.Println(a1)
	fmt.Printf("Length: %v\n", len(a1))
	fmt.Printf("Capacity: %v\n", cap(a1))
	b1 := a1  // The reference to slice a1 is passed to b1
	b1[1] = 5  // make a change to the same underlying copy of slice
	fmt.Println(a1)  // [1 5 3]
	fmt.Println(b1)  // [1 5 3]

	// More slice operations
	z := []int{1,2,3,4,5,6,7,8,9}
	a2 := z[:]  // slice of all elements
	b2 := z[3:]  // slice of 4th elements to end
	c2 := z[:5]  // slice first 5 elements
	d2 := z[2:5]  // slice the 3rd, 4th, 5th elements
	fmt.Println(z)
	fmt.Println(a2)
	fmt.Println(b2)
	fmt.Println(c2)
	fmt.Println(d2)
	// All the slices point to the same underlying array
	// slicing can also be operated on an array

	a3 := make([]int, 3, 5)
	fmt.Println(a3)
	fmt.Printf("Length: %v\n", len(a3))
	fmt.Printf("Capacity: %v\n", cap(a3))
	a3 = append(a3, 1)  // the second argument has to be of the same type as the slice
	a3 = append(a3, 2,3,4,5,6)
	a3 = append(a3, []int{7,8,9}...)  // a little trick
	fmt.Printf("Length: %v\n", len(a3))
	fmt.Printf("Capacity: %v\n", cap(a3))

	// concatenate two slices
	a4 := append(a3[:5], a3[7:]...) // Warning: this may change values in the underlying array!!!
	fmt.Println(a4)
}