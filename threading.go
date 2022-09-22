package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mem1 sync.Mutex
var mem2 sync.Mutex

type numStream interface {
	numStream()
}

type num1 struct {
	a [100]int
}

type num2 struct {
	b [100]int
}

func (x num1) numStream() {
	mem1.Lock()
	for i := range x.a {
		x.a[i] = i
	}

	for i := range x.a {
		time.Sleep(time.Millisecond)
		fmt.Printf("%d", x.a[i])
	}
	mem1.Unlock()
	wg.Done()
}

func (y num2) numStream() {
	mem2.Lock()
	for j := range y.b {
		y.b[j] = j
	}

	for j := range y.b {
		time.Sleep(time.Millisecond)
		fmt.Printf("%d", y.b[j])
	}
	mem2.Unlock()
	wg.Done()
}

func main() {
	var arr1 num1
	var arr2 num2

	wg.Add(2)
	go arr1.numStream()
	go arr2.numStream()
	wg.Wait()
}
