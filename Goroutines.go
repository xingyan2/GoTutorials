package main
import (
	"fmt"
	"runtime"
	"sync"
	//"time"
)
/**
Creating goroutines
Synchronization(WaitGroups, Mutexes)
Parallelism

Best practices:
	Don't create goroutines in libraries
	Let consumer control concurrency
	When creating a goroutine, know how it will end.
		Avoid subtle memory leaks
	Check for race conditions at compile time
 */

 var wg = sync.WaitGroup{}
 var counter = 0
 var m = sync.RWMutex{}

 func main() {
 	/** Example 1
 	var msg = "Hello"
 	go func(msg string) {  // pass the variable in case it gets modified
 		fmt.Println(msg)
	}(msg)
 	msg = "Goodbye"
 	time.Sleep(100 * time.Millisecond)  // this is a bad practice
 	*/

 	/** Example 2
	 var msg = "Hello"
	 wg.Add(1)
	 go func(msg string) {  // pass the variable in case it gets modified
		 fmt.Println(msg)
		 wg.Done()
	 }(msg)
	 msg = "Goodbye"
	 wg.Wait()
 	*/

 	runtime.GOMAXPROCS(100)
 	for i := 0; i < 10; i++ {
 		wg.Add(2)
 		m.RLock()
 		go sayHello()
 		m.Lock()
 		go increment()
	}
 	wg.Wait()
 }

 func sayHello() {
 	fmt.Printf("Hello #%v\n", counter)
 	m.RUnlock()
	wg.Done()
 }
 func increment() {
 	counter++
 	m.Unlock()
 	wg.Done()
 }