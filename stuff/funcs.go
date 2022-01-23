package stuff

import "fmt"

func Functions() {
	sum(1, 2, 3, 4, 5)

	fmt.Println("5 minus 3 is", subtract(5, 3))
	fmt.Println("5 plus 3 is", sum2(5, 3))

	fmt.Println("numbers plus each other is", *pointSum(3, 4, 5, 6))
	dividend, err := divide(5, 3)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(dividend)

	// declaring an anonymous function
	// anonymous functions are invoked asynchronously
	func() {
		fmt.Println("hello")
		// The last parentheses indicates that the func is being invoked
	}()

	// anonymous function can also be set to a variable
	f := func() {
		fmt.Println("This function is assigned to a variable")
		// when an anonymous function is assigned to a variable it doesn't need
		// the '()' after the closing brace
	}

	f()

	// Methods
	g := greeter{
		greeting: "Happy Eid",
		name: "...Ali",
	}
	g.greet()
	fmt.Println(g)
}

// This func accepts a series of values that will be wrapped into a slice
// The variadic param need to be the last param
func sum(values ...int) {
	result := 0
	for _, v := range values {
		result += v
	}

	fmt.Println("This is the result", result)
}

// The result doesn't need to be declared now since its declared in the
// header
func sum2(values ...int) (result int) {
	for _, v := range values {
		result += v
	}
	return
}

// This func returns a value unlike the sum func
func subtract(x, y int) int {
	result := x - y
	return result
}

// using a return type as a pointer and that returns the memory address
// of the result var
func pointSum(values ...int) *int {
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}

// returning multiple values and one of them being an error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

// working with methods
/*
	a method is the same as a func with just some extra syntax, in the
	following case the 'g greater' is the receiver
*/
type greeter struct {
	greeting string
	name     string
}

// g greeter is the value receiver 	
func (g *greeter) greet() {
	fmt.Println(&g)
	fmt.Println(g.greeting,g.name)
	g.name="...Everyone"
}
