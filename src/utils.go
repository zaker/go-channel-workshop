package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// START OMIT
func Do(work func(int), times int) func(int) bool {
	i := 0
	return func(ms int) bool {
		i++
		t1 := time.Now()

		if i > times {
			return false
		}
		fmt.Println(i, "work started", t1)
		work(ms)
		fmt.Println(i, "work took", time.Now().Sub(t1))
		return true

	}
}

func Something(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
	return
}

// END OMIT

func DoChan(work func(int), times int) func(int) bool {
	iChan := make(chan int, times)
	go func() {
		for i := 1; i <= times; i++ {
			iChan <- i
		}
		close(iChan)
	}()
	wg := &sync.WaitGroup{}
	wg.Add(times)

	//START CHAN OMIT
	return func(ms int) bool {
		select {
		case i := <-iChan:

			if i == 0 {
				wg.Wait()
				return false
			}

			t1 := time.Now()
			fmt.Println(i, "work started", t1)
			work(ms)
			fmt.Println(i, "work took", time.Now().Sub(t1))
			wg.Done()
			return true
		default:
			return false
		}
	}
	// END CHAN OMIT
}

func RandInt(max int) int {
	return rand.Int() % max
}

func DoSafe(work func(int), times int) func(int) bool {
	iChan := make(chan int, times)
	go func() {
		for i := 1; i <= times; i++ {
			iChan <- i
		}
		close(iChan)
	}()
	wg := &sync.WaitGroup{}
	m := &sync.Mutex{}
	wg.Add(times)

	//START SAFE OMIT
	return func(ms int) bool {
		select {
		case i := <-iChan:

			if i == 0 {
				wg.Wait()
				return false
			}

			m.Lock()
			t1 := time.Now()
			fmt.Println(i, "work started", t1)
			work(ms)
			fmt.Println(i, "work took", time.Now().Sub(t1))
			m.Unlock()
			wg.Done()
			return true
		default:
			return false
		}
	}
	// END SAFE OMIT
}
