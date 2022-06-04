package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int,10)
	go producer(ch)
	read(ch)
}

func producer(ch chan<- int) {
	for i := 0; i < 20; i++ {
		ch <- i
		time.Sleep(time.Millisecond*50)
	}
	close(ch)
}

func read(ch <-chan int) {
	for num := range ch {
		time.Sleep(time.Millisecond*100)
		fmt.Println(num)
	}
}
