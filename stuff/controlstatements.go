package stuff

import "fmt"

func Ifstatement(){

	if true{
		fmt.Println("This is true")
	}

	// less then
	num1 := 0 
	if num1 < 1{
		fmt.Println("This is true for num1")
	}

	// less then or equal to
	num2 := 1
	if num2 <= 1{
		fmt.Println("This is true for num2")
	}

	// greater then
	num3 := 15
	if num3 > 1{
		fmt.Println("This is true for num3")
	}

	// greater then or equal to
	num4 := 2
	if num4 >= 1{
		fmt.Println("This is true for num4")
	}

	// equal, comparison operator 
	num5 := 1
	if num5 == 1{
		fmt.Println("This is true for num5")
	}

	// not equal to
	num6 := 5 
	if num6 != 1{
		fmt.Println("This is true for num6")
	}

	// or operator
	num7 := 3
	if num7 > 1 || num7 < 6{
		fmt.Println("This is true for num7")
	}

	// and operator 
	num8 := 8
	if num8 >= 1 && num8 <= 9{
		fmt.Println("This is true for num8")
	}
	 
}

func Switch(){
	switch 7{
	case 1 ,3, 8: 
		fmt.Println("One")
	case 2 , 6, 4:
		fmt.Println("two")
		// fallthrough can be used if you want the case to continue
		// even if the condition has already been met
		fallthrough
	default:
		fmt.Println("Whatever")	
	}

	// type switching switch case
	var i interface{} = 1
	switch i.(type){
		case int:
			fmt.Println("This is an int")
			//This will break out of the case early
			break
			fmt.Println("This is the second thing that executes")
		case float32:
			fmt.Println("This is a float32")
		case string:
			fmt.Println("This is a string")
		default:
			fmt.Println("i is something else")
	}
}