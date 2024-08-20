package main

import (
	"fmt"
)


func main(){ // goroutine 1
	ch := make(chan string) // empty

	// goroutine 2
	go func(){
		ch <- "Full Cycle"
	}()

	msg := <-ch
	fmt.Println(msg)
}