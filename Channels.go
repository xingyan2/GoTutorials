package main

import (
	"fmt"
	"sync"
)

/**
	Channel basics, Restricting data flow, Buffered channels, Closing channels,
	For...range loops with channels, Select statements
 */

var wg = sync.WaitGroup{}

func main() {
	/** Example 1
	ch := make(chan int)
	wg.Add(2)
	go func() {
		i := <- ch
		fmt.Println(i)
		wg.Done()
	}()
	go func() {
		ch <- 42
		wg.Done()
	}()
	wg.Wait()
	*/

	/** Example 2
	ch := make(chan int)
	wg.Add(2)
	go func(ch <-chan int) {  // this is a receive-only channel
		i := <- ch
		fmt.Println(i)
		wg.Done()
	}(ch)
	go func(ch chan<- int) {  // this is a send-only channel
		ch <- 42
		wg.Done()
	}(ch)
	wg.Wait()
	*/

	ch := make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) {  // this is a receive-only channel
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int) {  // this is a send-only channel
		ch <- 42
		ch <- 27
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()

	// Extra info
	// Select statement and another channel can be used to signal the end of sender's job

	// done := make(chan struct{})  // this takes no memory space and works as a signal channel
	// when sender finishes, sender gives: "done <- struct{}{}"
}
