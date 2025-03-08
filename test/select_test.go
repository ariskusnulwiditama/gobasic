package test

import (
	"log"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(4 * time.Second)
		ch1 <-"Message from channel 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <-"Message from channel 2"
	}()

	for i:=0; i<2; i++ {
			select {
				case msg1 := <-ch1:
					log.Println("Received from channel 1:", msg1)
				case msg2 := <-ch2:
					log.Println("Received from channel 2:", msg2)
			}
	}
}