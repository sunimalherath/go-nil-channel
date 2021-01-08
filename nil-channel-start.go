package main

import (
	"fmt"
	"math/rand"
	"time"
)

func reader(ch chan int) {
	t := time.NewTimer(10 * time.Second)

	for {
		select {
		case i := <-ch:
			fmt.Printf("%d\n", i)
		case <-t.C:
			ch = nil
		}
	}
}

func writer(ch chan int) {
	stopper := time.NewTimer(2 * time.Second)
	starter := time.NewTimer(5 * time.Second)

	savedCh := ch

	for {
		select {
		case ch <- rand.Intn(42):
		case <-stopper.C:
			ch = nil
		case <- starter.C:
			ch = savedCh
		}
	}
}
func main() {
	ch := make(chan int) // unbuffered channel

	go reader(ch)
	go writer(ch)

	time.Sleep(20 * time.Second)
}
