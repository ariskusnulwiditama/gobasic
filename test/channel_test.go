package test

import (
	"fmt"
	"testing"
	"time"
)

func getAverage(numbers []int, ch chan float64) {
	sum := 0
	for _, v := range numbers {
		sum += v
	}

	ch <- float64(sum) / float64(len(numbers))
}

func getMax(numbers []int, ch chan int) {
	max := numbers[0]

	for _, v := range numbers {
		if max < v {
			max = v
		}
	}

	ch <- max
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <-x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func TestChan1(t *testing.T) {
	numbers := []int{3, 5, 7, 9, 11}

	avgCh := make(chan float64)
	maxCh := make(chan int)

	go getAverage(numbers, avgCh)
	go getMax(numbers, maxCh)

	

	for range 2 {
		select {
			case avg := <-avgCh:
				t.Log("Received average:", avg)
			case max := <-maxCh:
				t.Log("Received max:", max)
		}
	}
}

func TestChan2(t *testing.T) {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i:=0; i<10; i++ {
			fmt.Println(<-c)
		}
		quit <-0
	}()
	fibonacci(c, quit)
}

func TestChan3(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	channel3 := make(chan string)
	
	go workerer(1, channel1)
	go workerer(2, channel2)
	go workerer(3, channel3)

	time.Sleep(1 * time.Second)
	channel1 <- "Task for worker 1"
	channel2 <- "Task for worker 2"
	channel3 <- "Task for worker 3"

	close(channel1)
	close(channel2)
	close(channel3)

	time.Sleep(1 * time.Second)
}

func workerer(id int, ch chan string) {
	for v := range ch {
		fmt.Printf("Worker %d received %s\n", id, v)
	}
}

func TestChan4(t *testing.T) {
	messages := make(chan int, 4)

	go func ()  {
		for {
			i := <-messages
			fmt.Println("received data", i)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	for i := 1; i <= 5; i++ {
		fmt.Println("sending data", i)
		messages <- i
	}

	time.Sleep(2 * time.Second)
}