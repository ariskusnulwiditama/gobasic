package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d: Started\n", id)
	fmt.Printf("Worker %d: Finished\n", id)
}

func DownloadData(source string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Starting Download Data from %s\n", source)
	// simulate download data
	time.Sleep(time.Duration(len(source)) * time.Second)
	fmt.Printf("Finished Download Data from %s\n", source)
}

func Increment(counter *int, wg *sync.WaitGroup, mtx *sync.Mutex) {
	defer wg.Done()

	for i:=0; i<1000; i++ {
		mtx.Lock()
		*counter++
		mtx.Unlock()
	}
}

func TestWorker(t *testing.T) {
	wg := sync.WaitGroup{}

	for i:=0; i<=3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("All Workers Done")
}

func TestDownloadData(t *testing.T) {
	wg := sync.WaitGroup{}

	sources := []string{"Source A", "Source B", "Source C", "Source D"}
	for _, source := range sources {
		wg.Add(1)
		go DownloadData(source, &wg)
	}

	wg.Wait()
	fmt.Println("All Download Data Done")
}

func TestIncrement(t *testing.T) {
	wg := sync.WaitGroup{}
	mtx := sync.Mutex{}

	counter := 0

	for i:=0; i<5; i++ {
		wg.Add(1)
		go Increment(&counter, &wg, &mtx)
	}
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}