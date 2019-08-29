package main

import (
	"fmt"
	"local/utils"
	"time"
)

func main() {
	t1 := time.Now()
	// START OMIT
	ticker := time.NewTicker(time.Second)
	doSomethingFiveTimes := utils.Do(utils.Something, 5)
	for range ticker.C {

		cont := doSomethingFiveTimes(500)
		if !cont {
			break
		}
	}
	// END OMIT
	fmt.Println("It took a total of", time.Now().Sub(t1))
}
