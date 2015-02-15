package main

import (
    "github.com/adarqui/timeit-go"
    "time"
)


func nothingFunc() {
}


func loopFunc() {
    for i := 1; i < 1000000000; i++ {
    }
}


func sleepFunc() {
    time.Sleep(5 * time.Second)
}


func possibleError() (interface{}, interface{}) {
    return true, nil
}


func main() {
    timeit.TimeItPrint(nothingFunc)
    timeit.TimeItPrint(loopFunc)
    timeit.TimeItPrint(sleepFunc)
    timeit.TimeIt2Print(possibleError)
}
