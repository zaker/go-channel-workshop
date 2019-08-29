package main

import (
	"fmt"
	"local/utils"
	"sync"
	"time"
)

func main() {

	cont := make(chan bool, 1)
	t1 := time.Now()
	wg := &sync.WaitGroup{}
	nTimes := 5
	wg.Add(5)
	// START OMIT
	done := make(chan bool, 0)
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		doSomethingNTimes := utils.Do(utils.Something, nTimes)
		for {
			select {
			case <-ticker.C:
				cont <- doSomethingNTimes(utils.RandInt(3000))

			case <-done:
				return
			}

		}
	}()

	// END OMIT
	go func() {
		for c := range cont {
			if !c {
				done <- true
			} else {
				wg.Done()
			}
		}
	}()

	wg.Wait()
	fmt.Println("It took a total of", time.Now().Sub(t1))
}
