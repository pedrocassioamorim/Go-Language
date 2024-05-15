package main

import (
	"fmt"
	"time"
)

func worker(workerID int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d \n", workerID, x)
		time.Sleep(time.Second)
	}
}

func main() { //T1
	canal := make(chan int)
	qtdWorkers := 1000
	for i := 0; i < qtdWorkers; i++ {
		go worker(i, canal)
	}
	for i := 0; i < 10000; i++ {
		canal <- i
	}
}
