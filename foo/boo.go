package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func service1(c chan string) {
	c <- "Hello from service 1"
}

func service2(s string, c chan string) {
	c <- "**" + s
}

func mains() {

	chan1 := make(chan string)
	chan2 := make(chan string)

	go service1(chan1)
	go service2(getData(chan1, 2), chan2)
	response := getData(chan2, 2)
	fmt.Println(" Response ", response)
}

func getData(c chan string, timeoutSeconds int) string {
	select {
	case res := <-c:
		return res
	case <-time.After(time.Duration(timeoutSeconds) * time.Second):
		return ""
	}
}
