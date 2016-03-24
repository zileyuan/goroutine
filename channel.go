package main

import (
	"fmt"
"time"
)

var die chan bool

func worker(die chan bool, index int) {
	timeout := time.After(2 * time.Second)
	fmt.Println("Begin: This is Worker", index)
	for {
		select {
		case <-die:
			fmt.Println("Done1: This is Worker", index)
			die <- true
			fmt.Println("Done2: This is Worker", index)
			return
		//做事的分支
		case <- timeout:
			fmt.Println("Timeout", index)
			die <- true
			return
		}
	}
}

func doit(index int) {
	go worker(die, index)
	die <- true
	fmt.Println("Worker", index)
	<-die
}

func init() {
	fmt.Println("init")
	die = make(chan bool)
}

func main() {
	for i := 0; i < 50; i++ {
		go doit(i)
	}
	time.Sleep(50 * time.Second)
	fmt.Println("Worker goroutine has been terminated")
}