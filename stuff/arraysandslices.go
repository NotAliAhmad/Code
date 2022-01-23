package stuff

import "fmt"

func Arrays(){
	// setting an array that can hold 3 ints
	grades := [...]int{95,86,99}
	fmt.Printf("Grades: %v\n",grades)

	//you can also set the amount and pass in data that way
	var people [3]string
	people[0] = "mom"
	people[1] = "dad"
	people[2] = "brother"
	fmt.Printf("Grades: %v\n",people)

	//declaring an array of arrays - identity matrix
	var peoples [3][3]string
	peoples[0] = [3]string{"mom","mommy","mumzy"}
	peoples[1] = [3]string{"dad","daddy","danzy"}
	peoples[2] = [3]string{"bro","bra","broder"}
	fmt.Printf("Grades: %v\n",peoples)


} 

func Slices(){
	/*
	Arrays create a new place in memory so they don't affect perviously
	altered versions.Slices are changed by reference therefore changing 
	it in one place will affect it in another.
	*/

	// this is an array
	// putting in 3 dots[...] or setting the sizing in the brackets makes it an array
	var a1 = [3]int{1,2,3}
	fmt.Println(a1)
	// this is a slice
	var a = []int{1,2,3}
	b := a
	b[1] = 10
	// be is the same as a
	fmt.Printf("Grades: %v %T\n",b,b)
	fmt.Printf("Grades: %v %T\n",a,a)

	// another way to set a slice
	// first arg is the type of obj, 
	// second arg is the length of the slice
	// the third arg is the capacity 
	ah := make([]int,0)
	fmt.Println(ah)
	fmt.Println(cap(ah))

	// I can add to my slice with the append method
	ah = append(ah, 32)
	fmt.Println(cap(ah))

	// I can use the append method to add multiple entries at once
	ah = append(ah, []int{54,23,99,21}...)
	fmt.Println(ah)
	// cap() returns the value of the underlying array since a slice are backed by an array
	fmt.Println(cap(ah))

	// removing an entry from the beginning of the slice
	bh := ah[1:]
	fmt.Println(bh)

	// removing an entry from the end of the slice
	ch := ah[:len(ah)-1]
	fmt.Println(ch)

	// removing an entry from the middle of the slice
	// includes entry 4 but not entry 3
	dh := append(ah[:3],ah[4:]...)
	fmt.Println(dh)

	// because we're using a slice removing the value from the old slice in the new slice affects the old one
	fmt.Println(ah)


}