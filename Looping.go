package main
import (
	"fmt"
)
/**
For statements
simple loops, exit early, looping through collections
 */
func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	j := 0
	for ; j < 5; {  // semicolons are required
		fmt.Println(j)
		j++
		if j == 3 {
			break  // exit the loop
		}
	}

	// break the nested for loop
	// add a label
Loop:
	for i := 1; i < 3; i++ {
		for j:= 1; j < 3; j++ {
			if i * j > 4 {
				break Loop
			}
			fmt.Println(i * j)
		}
	}

	// Work with collection variables
	// slice, array, map, string
	s := []int{1,2,3,4}
	for k, v := range s {
		fmt.Println(k, v)
	}
	for _, v := range s { // ignore the key
		fmt.Println(v)
	}
	for k := range s { // ignore the value
		fmt.Println(k)
	}
}