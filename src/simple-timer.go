package main

import (
	"fmt"
	"local/utils"
	"time"
)

func main() {
	t1 := time.Now()
	// START OMIT
	doSomethingFiveTimes := utils.Do(utils.Something, 5)
	dur := 1 * time.Second
	for {
		time.Sleep(dur)
		cont := doSomethingFiveTimes(500)

		if !cont {
			break
		}
	}
	// END OMIT
	fmt.Println("It took a total of", time.Now().Sub(t1))
}
