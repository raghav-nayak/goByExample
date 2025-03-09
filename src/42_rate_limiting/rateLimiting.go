// Rate limiting is an important mechanism for controlling resource utilization and maintaining quality of service
// Go elegantly supports rate limiting with goroutines, channels and tickers

package main

import (
	"fmt"
	"time"
)

func main() {
    requests := make(chan int, 5) // buffered channel
    for i := 1; i <= 5; i++ {
        requests <- i
    }
    close(requests)

    // limiter channel will receive a value every 200 milliseconds
    // This is the regulator in our rate limiting scheme
    limiter := time.Tick(200 * time.Millisecond)

    for req := range requests {
        // blocking on a receive from the limiter channel before serving each request
        // this limits to 1 request every 200 milliseconds
        <-limiter 
        fmt.Println("request", req, time.Now())
    }

    // to allow short bursts of requests in this rate limiting scheme while preserving the overall rate limit
    // this is done using buffered channel; allows bursts of up to 3 events
    burstyLimiter := make(chan time.Time, 3)

    // fill up the channel represent allowed bursting
    for i := 0; i <3; i++ {
        burstyLimiter <- time.Now()
    }

    // every 200 milliseconds, we will try to add a new value to burstyLimiter, up to its limit of 3
    go func ()  {
        for t := range time.Tick(200 * time.Millisecond) {
            burstyLimiter <- t
        }
    }()

    // simulate 5 more incoming requests
    // the first 2 of these will benefit from the burst capability of burstyLimiter
    burstyRequests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        burstyRequests <- i
    }
    close(burstyRequests)
    for req := range burstyRequests {
        <-burstyLimiter
        fmt.Println("request", req, time.Now())
    }
}

// output:
//// first batch of requests handled once every ~200 milliseconds as desired
// request 1 2025-03-10 00:06:12.325462362 +0530 IST m=+0.200480301
// request 2 2025-03-10 00:06:12.52579627 +0530 IST m=+0.400814145
// request 3 2025-03-10 00:06:12.725214451 +0530 IST m=+0.600232318
// request 4 2025-03-10 00:06:12.925623351 +0530 IST m=+0.800641212
// request 5 2025-03-10 00:06:13.126047867 +0530 IST m=+1.001065696
//// second batch of requests, the first 3 requests are served immediately because of burstable rate limiting, the remaining 2 with ~200ms delays each
// request 1 2025-03-10 00:06:13.126111456 +0530 IST m=+1.001129284 - burst
// request 2 2025-03-10 00:06:13.126120784 +0530 IST m=+1.001138603 - burst
// request 3 2025-03-10 00:06:13.126125097 +0530 IST m=+1.001142920 - burst
// request 4 2025-03-10 00:06:13.326482641 +0530 IST m=+1.201500477
// request 5 2025-03-10 00:06:13.526758597 +0530 IST m=+1.401776425
 