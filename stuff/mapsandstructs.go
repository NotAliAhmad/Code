package stuff

import (
	"fmt"
	"reflect"
)

func Maps() {
	/*
		Like slices, altering a value in the original map and establing it
		again will affect both
	*/

	// Creating a map can be done if the date is not predefined
	statePopulations := make(map[string]int)

	// If the data is defined, creating and setting it can be done in the same line
	statePopulations = map[string]int{
		"California":   39543545,
		"Texas":        27849996,
		"Florida":      20456364,
		"New York":     19242444,
		"Pennsylvania": 12945365,
		"Illinois":     12845936,
		"Ohio":         11386947,
	}
	fmt.Println(statePopulations)

	// adding a value to the map
	statePopulations["Georgia"] = 10387474
	fmt.Println(statePopulations)

	// deleting a value in the map
	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations)

}

// This is how you create a type of struct
type Doctor struct {
	Number     int
	ActorName  string
	Companions []string
}

func Structs() {
	/*
		Like arrays structs that are duplicate will have independent values
	*/
	// initializing a struct
	aDoctor := Doctor{
		Number:    3,
		ActorName: "Jon Snow",
		Companions: []string{
			"Leo Decapri",
			"Jessica Alba",
			"Morgan Freeman",
		},
	}
	fmt.Println(aDoctor)

	// This is an anonymous struct
	bDoctor := struct{ name string }{name: "Jay Gerrik"}

	fmt.Println(bDoctor)

}

type Animal struct {
	Name   string `required max:"100"`
	Origin string
}

// Embedding a struct
// There is no inheritance in go but we can use composition thru embedding
type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func Embedding(){
	
	/*
	This is individually setting each field of the bird struct
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	*/

	// This is the literal syntax  
	b := Bird{
		Animal: Animal{Name:"Emu",Origin:"Australia"},
		SpeedKPH: 30,
		CanFly: false,
	}
	fmt.Println(b)

	// Getting a tag from a field
	t := reflect.TypeOf(Animal{})
	field,_ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}