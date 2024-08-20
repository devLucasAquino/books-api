package main

import (
	"fmt"
	"time"
)

func worker(workerID int, data chan int){
	for x := range data{
		fmt.Printf("worker %d got %d\n", workerID, x)
		time.Sleep(time.Second)
	}
}


func main(){ // goroutine 1
	ch := make(chan int) // empty
	qntWorkers := 10

	for i:= range qntWorkers{
		go worker(i, ch)
	}
	
	for i := range 10{
		ch <- i 
	}
}