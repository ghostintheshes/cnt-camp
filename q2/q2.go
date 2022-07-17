package q2

import (
	"fmt"
	"math/rand"
	"time"
)

func Producer(ch chan<- int) {
	for {
		ch <- rand.Intn(10)
		time.Sleep(time.Second * 1)
	}
}

func Consumer(ch <-chan int) {
	for true {
		time.Sleep(time.Second * 1)
		n := <-ch
		fmt.Println(n)
	}
}
