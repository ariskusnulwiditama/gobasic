package test

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <-"Message from channel 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <-"Message from channel 2"
	}()

	for {
		select {
			case msg1 := <-ch1:
				log.Println("Received from channel 1:", msg1)
				/* return */
			case msg2 := <-ch2:
				log.Println("Received from channel 2:", msg2)
				// return
		}
	}
}

func TestSelectChannel(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Simulasi layanan yang butuh waktu berbeda
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Respons dari Layanan 1"
	}()
	go func() {
		time.Sleep(5 * time.Second)
		ch2 <- "Respons dari Layanan 2"
	}()

	// select akan menunggu sampai salah satu channel siap
	select {
	case msg1 := <-ch1:
		fmt.Println("Menerima:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Menerima:", msg2) // Ini yang akan jalan karena lebih cepat
	}
}