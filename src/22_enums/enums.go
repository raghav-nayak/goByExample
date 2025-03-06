// Enumerated types(enums) are a special case of sum types(abstract data types)
// An enum is a type that has a fixed number of possible values, each with a distinct name
// read about go generate : https://eli.thegreenplace.net/2021/a-comprehensive-guide-to-go-generate

package main

import "fmt"

type ServerState int

// possible values are defined as constants
const (
    StateIdle ServerState = iota // iota generates successive constant value automatically; from 0,1,2 so on.
    StateConnected
    StateError
    StateRetrying
)

// this implements fmt.Stringer interface
var stateName = map[ServerState]string {
    StateIdle: "idle",
    StateConnected: "connected",
    StateError: "error",
    StateRetrying: "retrying",
}

func (ss ServerState) String() string {
    return stateName[ss]
}

func main() {
    ns1 := transition(StateIdle) // we cannot pass int here; for compile time safety
    fmt.Println(ns1)
    
    ns2 := transition(ns1) // passing the previous state
    fmt.Println(ns2)
}

func transition(s ServerState) ServerState {
    switch s {
    case StateIdle:
        return StateConnected
    case StateConnected, StateRetrying:
        return StateIdle
    case StateError:
        return StateError
    default:
        panic(fmt.Errorf("unknown state: %s", s))
    }
}


// output:
// connected
// idle