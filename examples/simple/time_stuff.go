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
	timeit.TimeitPrint(nothingFunc)
	timeit.TimeitPrint(loopFunc)
	timeit.TimeitPrint(sleepFunc)
	timeit.Timeit2Print(possibleError)
}
