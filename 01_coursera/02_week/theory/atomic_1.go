package main

import (
	"fmt"
	"time"
)

func main() {
	var totalOperations int32 = 0

	inc := func() {
		totalOperations++
	}

	for i := 0; i < 1000; i++ {
		go inc()
	}
	time.Sleep(2 * time.Millisecond)
	// ожидается 1000, но по факту будет меньше
	fmt.Println("total operation = ", totalOperations)
}
