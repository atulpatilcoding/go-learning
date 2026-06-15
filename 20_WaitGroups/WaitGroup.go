package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done() // Decrement the WaitGroup counter when the task is done
	fmt.Println("Doing Task ", id)
}

func main() {
	var wg sync.WaitGroup // Create a WaitGroup to wait for all tasks to complete there is sync package for WaitGroup
	for i := 0; i < 11; i++ {
		wg.Add(1) // Increment the WaitGroup counter for each task	
		go task(i, &wg) // Start a goroutine for each task and pass the WaitGroup
	}

}
