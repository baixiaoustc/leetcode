package leetcode

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test71(t *testing.T) {
	ch1, ch2 := make(chan int), make(chan int)
	go func(ch chan int) {
		for i := 1; i <= 99; i += 2 {
			ch <- i
			time.Sleep(time.Millisecond)
		}
		close(ch)
	}(ch1)
	go func(ch chan int) {
		for i := 2; i <= 100; i += 2 {
			ch <- i
			time.Sleep(time.Millisecond)
		}
		close(ch)
	}(ch2)
	for {
		select {
		case c := <-ch1:
			fmt.Println(c)
		case c := <-ch2:
			fmt.Println(c)
		}
	}
}

func Test72(t *testing.T) {
	ch := make(chan int)
	mtx := new(sync.Mutex)
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func() {
		for i := 1; i <= 99; i += 2 {
			mtx.Lock()
			ch <- i
			mtx.Unlock()
			time.Sleep(time.Duration(i) * time.Microsecond)
		}
		wg.Done()
	}()
	time.Sleep(10 * time.Millisecond)
	go func() {
		for i := 2; i <= 100; i += 2 {
			mtx.Lock()
			ch <- i
			mtx.Unlock()
			time.Sleep(time.Duration(i) * time.Microsecond)
		}
		wg.Done()
	}()
	go func() {
		wg.Wait()
		close(ch)
	}()
	for c := range ch {
		fmt.Println(c)
	}
}
