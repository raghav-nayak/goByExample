// Tickers are for when you want to do something repeatedly at regular intervals.
// Similar to Timers

package main

import (
	"fmt"
	"time"
)

func main() {
    // a channel that is sent values
	ticker := time.NewTicker(500 * time.Millisecond) // to await the value as they arrive every 500ms
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C: 
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(2000 * time.Millisecond) // we will stop after 2s; if you remove this, we won't get any output
    ticker.Stop() // similar to Timers; stops the ticker
    done <- true
    fmt.Println("Ticker stopped")
}

// output:
// note : 4 lines as 500ms * 4 = 2s
// Tick at 2025-03-09 21:57:18.452597241 +0530 IST m=+0.500036749
// Tick at 2025-03-09 21:57:18.952597 +0530 IST m=+1.000036499
// Tick at 2025-03-09 21:57:19.452597673 +0530 IST m=+1.500037161
// Tick at 2025-03-09 21:57:19.952596282 +0530 IST m=+2.000035762
// Ticker stopped