// Basic sends and receives on channels are blocking.
// We can use select with a default clause to implement non-blocking sends, receives and
// non-blocking multi-way selects.
// TODO read more about this

package main

import "fmt"

func main() {
	messages := make(chan string) // channel without buffer
	signals := make(chan bool)

	// non-blocking receive
	select {
	case msg := <-messages: // if a value is available on messages then select will take the value
		fmt.Println("received message", msg)
	default: // if there is value available, it will take the default case
		fmt.Println("no message received") // this will be printed
	}

	// non-blocking send
	msg := "hi"
	select {
	case messages <- msg: // msg cannot be sent to the channel without buffer and there is no receiver
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent") // this will be printed
	}

	// multi-way non-blocking select
	select {
	case msg := <-messages: // non-blocking receive
		fmt.Println("received message", msg)
	case sig := <-signals: // non-blocking receive
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity") // this will be printed
	}
}

// output:
// no message received
// no message sent
// no activity
