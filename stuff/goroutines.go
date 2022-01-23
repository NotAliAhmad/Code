package stuff

import (
	"fmt"
	"runtime"
	"sync"
)
// When running go routines, the go run - race main.go can help detect bugs
func Routines(){
	/*
		Go Routines are an abstraction of a thread
		it is common to use a func to invoke a go routine 

		The key word go is used to indicate that a go routine is being used
	*/
	var msg = "Run Yo SHEET go Routine !!!"
	wg.Add(1)
	// The keyword go is used to indicate that a go routine is being used
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)
	msg = "GoodBye"
	wg.Wait() 


	// When you want to fire off 2 routines in sync the sync package is 
	// useful for that
}

var wg sync.WaitGroup
var counter = 0
var m = sync.RWMutex{}


func Mutex(){
	// Mutex is a lock that the app is going to honor
	// I am using th mutex to lock the code to run asynconesly 
	// this is how many cores are available to the program 
	runtime.GOMAXPROCS(100)
	fmt.Println("Threads:",runtime.GOMAXPROCS(-1))
	for i := 0; i < 10; i++{
		wg.Add(2)
		// This is a read lock
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello(){
	fmt.Println("Right now the number is ",counter)
	m.RUnlock()
	wg.Done()
}

func increment(){
	counter++
	m.Unlock()
	wg.Done()
}


