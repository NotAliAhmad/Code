package stuff

import "fmt"

func Loops(){
	
	// Regular for loop
	for i:=0;i<5;i++{
		fmt.Println(i)
	}

	// Nested for loop
// This is my label to where the loop will break out to	
Loop:	
	for i:=0;i<1;i++{
		fmt.Println("This is i",i)
		for j:=0;j<5;j++{
			fmt.Println("This is j",j)
			/*
				The regular break will only break out of the j loop but
				putting a label(I called mine 'Loop') will break out to where
				the label tag is
			*/
			break Loop
		}
	}

	// For range loop
	s := []int{1,2,3}
	for k, v := range s{
		fmt.Println("This is the key",k)
		fmt.Println("This is the ~~~~~ value",v)
	}

	// Range over a map
	statePopulations := map[string]int{
		"California":   39543545,
		"Texas":        27849996,
		"Florida":      20456364,
		"New York":     19242444,
		"Pennsylvania": 12945365,
		"Illinois":     12845936,
		"Ohio":         11386947,
	}

	for k, v := range statePopulations{
		fmt.Println(k,"has",v,"people")
	}

	// Ranging over a string
	b := "This is a string"
	for k,v := range b{
		// The value by default returns the unicode value and the 
		// string() method takes that and converts it into a string
		fmt.Println("This char is ",string(v),"and it is in the index",k)
	}
}