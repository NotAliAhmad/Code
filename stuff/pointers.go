package stuff

import "fmt"

type myStruct struct {
	foo int
}

func Point() {
	a := 42
	// This is a pointer
	b := &a
	// The star * is the dereferencing operatoring which looks inside
	// that memory address and gives back the value stored in it
	fmt.Println(&a, *b)

	// Nil are values that aren't initialized
	var ms *myStruct
	fmt.Println(ms)

	// new() zeros out all the values and does the same as '&'
	ms = new(myStruct)
	fmt.Println(ms)

	// to work with an uninitialized variable we need to set a value
	ms.foo = 40
	fmt.Println(ms.foo)

}
