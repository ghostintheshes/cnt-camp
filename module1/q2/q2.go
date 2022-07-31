package q2

import (
	"fmt"
	"math/rand"
	"time"
)

func Producer(ch chan<- int) {
	for {
		n := rand.Intn(10)
		fmt.Printf("produce number: %d\n", n)
		ch <- n
		time.Sleep(time.Second * 1)
	}
}

func Consumer(ch <-chan int) {
	for true {
		time.Sleep(time.Second * 1)
		n := <-ch
		fmt.Println(fmt.Printf("consume a number: %d\n", n))
	}
}
