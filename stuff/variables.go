package stuff

import (
	"fmt"
	"strconv"
)

func VarDeclaration(){
	// one way
	var i int
	i = 42
	fmt.Printf("%v, %T\n",i,i)

	// second way
	var j float32 = 27.00
	fmt.Printf("%v, %T\n",j,j)

	// third way
	k := 99
	fmt.Printf("%v, %T\n",k,k)

}

func VarConverstion(){
	// regular int
	var i int = 42
	fmt.Printf("%v, %T\n",i,i)

	// converting a int to a string
	var k string
	k = strconv.Itoa(i)
	fmt.Printf("%v, %T\n",k,k)

	// converting a string to an int
	var a int 
	a,_ = strconv.Atoi(k)
	fmt.Printf("%v, %T\n",a,a)
}