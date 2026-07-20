package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	mychannel := make(chan string, 100)
	var wg sync.WaitGroup
	wg.Add(2)
	go pipeline(mychannel, &wg)
	go printlogs(mychannel, &wg)
	wg.Wait()
		
}

func Sleep() {
	time.Sleep(2 * time.Second)
}
func pipeline(mychannel chan string, wg *sync.WaitGroup) {
	//step 1
	mychannel <- "step 1"

	//step2
	mychannel <- "step2"
	Sleep()

	//step3
	mychannel <- "step3"

	//done
	mychannel <- "done"
	defer wg.Done()
	defer close(mychannel)
}

func printlogs(mychannel chan string, wg *sync.WaitGroup) {
	for log := range mychannel {
		fmt.Print(log)
	}
	defer wg.Done()
	
}