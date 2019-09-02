package main
import (
	"fmt"
	"strconv"
)

var l float32 = 15  // Declare the variable at package level, visible to any files within the package

var L int = 100  // uppercase variable declared at the package level is exported from the package and globally visible

var (
	actor string = "Tony"
	actress string = "Emma"
)

func main(){
	var i int	// Declare a new int variable i
	i = 12
	var j int = 13	// The data type is specified
	k := 14		// Let the compiler to figure out the data type
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Printf("%v, %T\n", j, j)

	// Check how the scope of variables exist
	fmt.Println(l)
	var l float32 = 28
	fmt.Println(l)

	// Convert from one data type to another
	// int to float32
	var m float32
	m = float32(i)
	fmt.Printf("%v, %T\n", m, m)

	// int to string
	var str string
	str = strconv.Itoa(k)
	fmt.Printf("%v, %T\n", str, str)
}




