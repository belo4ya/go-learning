package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var totalOperations int32 = 0

	inc := func() {
		atomic.AddInt32(&totalOperations, 1) // атомарно
	}

	for i := 0; i < 1000; i++ {
		go inc()
	}
	time.Sleep(2 * time.Millisecond)
	fmt.Println("total operation = ", totalOperations)
}
