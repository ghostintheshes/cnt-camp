package q2

import (
	"log"
	"testing"
	"time"
)

func TestConsumer(t *testing.T) {
	ch := make(chan int, 10)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Fatal(err)
			}
		}()
		Producer(ch)
	}()
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Fatal(err)
			}
		}()
		Producer(ch)
	}()
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Fatal(err)
			}
		}()
		Producer(ch)
	}()
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Fatal(err)
			}
		}()
		Consumer(ch)
	}()
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Fatal(err)
			}
		}()
		Consumer(ch)
	}()
	time.Sleep(time.Second * 11)
}
