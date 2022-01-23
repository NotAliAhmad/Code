package stuff

import (
	"fmt"
)

func Deferred(){
	/*
		The most common use case for defer is closing a resource since
		they need to be closed in the reverse order that they were opened
		in case there is some sort of dependence.
	*/

	defer fmt.Println("Start")
	// defer works by the Last in First out order for multiple defer calls
	// defers delay the execution of a function until the end of the program
	defer fmt.Println("Middle")
	fmt.Println("End")
}

func PanicAndRecover(){
	fmt.Println("Starting a program")
	/*
		Panics happen before the deferred statements are executed so 
		they can still run after the panic happens
	*/
	// panic("sheeeet, something failed")
	fmt.Println("Not calling panic due to it currently being unhandled")
	// fmt.Println("Ending program due to a panic being thrown")
	/*
		If panic is handled using recover() then it stops executing the 
		function but whatever calls the function can continue to run 
	*/

	/*
		Recover:
		- Used to recover from panics
		- Only useful in deferred funcs
		- Current function will not attempt to continue but higher functions
		in the call stack will
		- If the panic is not something that can be handled then panic can
		be called again in the anonymous func to rethrow that panic

	*/
}

