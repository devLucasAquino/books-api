package main

import (
	"fmt"
	"time"
)

func counter(n int) {
	for i := range n {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main(){
	go counter(5)
	go counter(5)
	counter(5)
}