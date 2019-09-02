package main
import (
	"fmt"
	"reflect"
)
/**
Maps: What are they? Creating and manipulating
Structs: What are they? Creating, Naming conventions, Embedding, and Tags
*/

type Doctor struct {
	number int
	actorName string
	companions []string
}

// Composition of structs
type Animal struct {
	Name string
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly bool
}

type Dog struct {
	Name string `required max: "100"`
	Origin string
}

func main() {
	// Create a map
	// statePopulations := make(map[string]int, 10)
	statePopulations := map[string]int{
		"California": 39250017,
		"Pennsylvania": 12802503,
		"New York": 19745289,
	}
	//fmt.Println(statePopulations)
	//copy := statePopulations // map is copy by reference
	fmt.Println(statePopulations["California"])
	statePopulations["Ohio"] = 11614373  // Add a new entry to the map
	fmt.Println(statePopulations)
	delete(statePopulations, "Ohio")  // delete from the map
	fmt.Println(statePopulations)
	// Get the length of the map
	fmt.Println(len(statePopulations))

	// Check if the key is found
	pop, ok := statePopulations["Michigan"]
	fmt.Println(pop, ok)

	// Structs
	aDoctor := Doctor{
		number: 3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
		},
	}

	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)

	//copy := aDoctor // struct is copy by value

	// Composition of structs via embedding
	// Bird has Animal characteristics
	bi := Bird{}
	bi.Name = "Emu"
	bi.Origin = "Australia"
	bi.SpeedKPH = 48
	bi.CanFly = false
	fmt.Println(bi)

	// Tags
	// Tags can be added to struct fields to describe field
	t := reflect.TypeOf(Dog{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
