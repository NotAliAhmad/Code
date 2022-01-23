package stuff

import "fmt"
/*
	declaring an interface
	Interfaces describe behavior more then data

	Implementing interfaces with values vs pointers
	- Method set of VALUE is all methods with value receivers
	- Method set of POINTERS is all methods, regardless of receiver type
	
	** if the inferfaces are implemented with pointers then I have
	to have the methods there regardless if the receiver has a pointer
	or not 
*/
type Writer interface{
	Write([]byte) (int, error)
}

type ConsoleWriter struct{
	
}
func Interfaces(){
	var w Writer = &ConsoleWriter{}
	w.Write([]byte("I AM INTERFACE"))  

}

func (cw *ConsoleWriter) Write(data[]byte)(int, error){
	// This takes in a set of bytes(Unicode values) and converts it into
	// a string
	n, err := fmt.Println(string(data))
	return n,err
}

/*
	When composing multiple interfaces:
	type Writer interface{
		Writer([]byte)(int,error)
	}

	type Close interface{
		Closer()(error)
	}

	type WriterCloser interface{
		Writer
		Closer
	}

	For this example, both the Writer and Closer methods need to be 
	implemented. And if this WriterCloser interface need to be passed
	around the application. It can be implemented as a new Writer
	interface. Ex. var wi Writer = WriterCloser{}
*/