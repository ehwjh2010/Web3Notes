package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter interface {
	Incr()
	Count() int64
}

type MutexCounter struct {
	mutex sync.Mutex
	count int64
}

func NewMutexCounter() Counter {
	return &MutexCounter{}
}

func (m *MutexCounter) Incr() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.count++
}

func (m *MutexCounter) Count() int64 {
	return m.count
}

func IncrCounter(m Counter) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				m.Incr()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(m.Count())
}

type AtomicCounter struct {
	count int64
}

func NewAtomicCounter() Counter {
	return &AtomicCounter{}
}

func (m *AtomicCounter) Incr() {
	atomic.AddInt64(&m.count, 1)
}

func (m *AtomicCounter) Count() int64 {
	return atomic.LoadInt64(&m.count)
}

func main() {
	mutexCounter := NewMutexCounter()
	IncrCounter(mutexCounter)

	atomicCounter := NewAtomicCounter()
	IncrCounter(atomicCounter)
}
