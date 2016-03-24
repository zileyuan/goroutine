package main

import (
"fmt"
"time"
"sync"
)

var lock *sync.Mutex

func worker1(index int) {
	fmt.Println("Begin: This is Worker", index)
	lock.Lock()
	defer lock.Unlock()
	fmt.Println("Done: This is Worker", index)
}

func doit1(index int) {
	go worker1(index)
	fmt.Println("Worker", index)
}

func init() {
	fmt.Println("init")
	lock = &sync.Mutex{}
}

func main() {
	for i := 0; i < 50; i++ {
		go doit1(i)
	}
	time.Sleep(50 * time.Second)
	fmt.Println("Worker goroutine has been terminated")
}