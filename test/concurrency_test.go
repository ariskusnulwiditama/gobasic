package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func SayHello(say string) {
	fmt.Printf("Hello %s \n", say)
}

func TestConcurrent(t *testing.T) {
	go SayHello("Golang")
	go SayHello("Goroutine")
	time.Sleep(2 * time.Second)
	t.Log("Finished")
}

// dengan waitgroup
func Worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	t.Log("All workers done")
}

// concurrency dengan channel
func WorkerChannel(id int, jobs <-chan int, result chan<- int) {
	for job := range jobs {
		result <- job * 2
	}
}

func TestChannelConcurrency(t *testing.T) {
	jobs := make(chan int, 5)
	result := make(chan int, 5)

	for i := 1; i <= 3; i++ {
		go WorkerChannel(i, jobs, result)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}

	close(jobs)
	
	// print results
	for a := 1; a <= 5; a++ {
		t.Log(<-result)
	}
}

// banyak channel dengan select

func WorkerChan1(ch1 chan<- string) {
	time.Sleep(2 * time.Second)
	ch1 <- "Data from channel 1"
}

func WorkerChan2(ch2 chan<- string) {
	time.Sleep(10 * time.Second)
	ch2 <- "Data from channel 2"
}

func TestSelectManyChannels(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go WorkerChan1(ch1)
	go WorkerChan2(ch2)
	
	for range 2 {
		select {
		case msg1 := <- ch1:
			t.Log(msg1)
		case msg2 := <- ch2:
			t.Log(msg2)
		}
	}
}

// worker pool

func WorkerPool(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(2 * time.Second)
		results <- fmt.Sprintf("Worker %d finished job %d", id, job)
	}	
}

func TestWorkerPool(t *testing.T) {
	jobs := make(chan int, 5)
	results := make(chan string, 5)
	var wg sync.WaitGroup

	// start 3 workers
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go WorkerPool(w, jobs, results, &wg)
	}

	// close channel jobs after all jobs are sent
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	
	// close results channel after all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// print results
	for res := range results {
		t.Log(res)
	}
}