package stuff

import (
	"fmt"
	"sync"
)

var wd sync.WaitGroup

func Channels(){
	// the make method must be used when creating a channel
	ch := make(chan int) // This channel can only send ints
	wd.Add(2)

	// This is the receiving go routine and a receiving func
	go func(ch <-chan int){
		i := <- ch
		fmt.Println(i)
		wd.Done()
	}(ch)

	// This is the sending go routine and a sending func
	go func(ch chan<- int){
		// The sending data is a copy so changing it here wont effect other places
		ch <- 42 
		wd.Done()
	}(ch)
	wd.Wait()
}

func RangeChannel(){
	/* 
		The 6 as the second param can decouple sender and receiver 
		with buffer channels; Buffer channels are good if you have more sending 
		data than receiveing data
	*/
	ch := make(chan int,6) // This channel can only send ints
	wd.Add(2)

	// This is the receiving go routine and a receiving func
	go func(ch <-chan int){
// A single channel is handling multiple data sends from a single channel
		for i := range ch{ // i is the value itself here, not the index
			fmt.Println(i)
		}
		wd.Done()
	}(ch)

	// This is the sending go routine and a sending func
	go func(ch chan<- int){
		// Sending in multiple data to a single channel
		ch <- 82 
		ch <- 99
		// Once data is finished being transmitted, the channel needs to be closed
		close(ch)
		wd.Done()
	}(ch)
	wd.Wait()
}

/* 
	A good way to monitor channels is to use a select case where it will
	close out the channel once data is send to the close logger which will break 
	break out of the select statement 

	for example;

	func logger(){
		for{
			select{
			case entry := <- logCh"
				fmt.Println("The channel is open",action)
			case <-doneCh:
				break
			}
		}
	}
 */