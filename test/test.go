package main

import (
	"fmt"
	"github.com/adskyfly/queue4go"
	// "math/rand"
	// "runtime"
	"sync"
	"time"
)

type mytype struct {
	name string
	age  int
}

func main() {
	// runtime.GOMAXPROCS(1)
	queue := queue4go.Queue("test")
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(5)
	queue.SetMaxLength(2)
	// fmt.Println(queue.GetMaxLength())
	// rand.Seed(time.Now().UnixNano())
	go func() {
		defer wg.Done()
		for i := 0; i < 1; i++ {
			// queue.Push(1)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 1; i++ {
			// queue.Push(2)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 1; i++ {
			queue.Push(3)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 1; i++ {
			queue.Push(4)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 1; i++ {
			queue.Push(5)
		}
	}()
	wg.Wait()
	delta := time.Since(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
	ql := queue.Length()
	for i := 0; i < ql; i++ {
		fmt.Println(queue.Pop())
	}
}
