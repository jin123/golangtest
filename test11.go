package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	t := int64(0)
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println("bef_i=", i)
			wg.Done()
			fmt.Println("after_i=", i)
			//t = 100
			atomic.AddInt64(&t, 100)
		}(i)
	}

	wg.Wait()

	fmt.Println("t=", t)
}
