package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	t1 := time.Now()
	nTimes := 5
	wg := &sync.WaitGroup{}

	// START OMIT
	for i := 0; i < 1000; i++ {
		wg.Add(nTimes)
		go func() {
			cont := make(chan bool, 1)
			ticker := time.NewTicker(time.Second)
			defer ticker.Stop()
			doSomethingNTimes := DoSafe(Something, nTimes)
			for {
				select {
				case <-ticker.C:
					go func() {
						cont <- doSomethingNTimes(RandInt(3000))
					}()
				case c := <-cont:
					if !c {
						return
					}
					wg.Done()
				}
			}
		}()
	}
	// END OMIT
	wg.Wait()
	fmt.Println("It took a total of", time.Now().Sub(t1))
}
