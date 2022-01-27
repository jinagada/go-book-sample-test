package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

type counter struct {
	i int64
}

func (c *counter) increment() {
	c.i += 1
}

func (c *counter) display() {
	fmt.Println(c.i)
}

func main1() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := counter{i: 0}
	done := make(chan struct{})
	for i := 0; i < 1000; i++ {
		go func() {
			c.increment()
			done <- struct{}{}
		}()
	}
	for i := 0; i < 1000; i++ {
		<-done
	}
	c.display()
}

type counter2 struct {
	i  int64
	mu sync.Mutex
}

func (c *counter2) increment() {
	c.mu.Lock()
	c.i += 1
	c.mu.Unlock()
}

func (c *counter2) display() {
	fmt.Println(c.i)
}

func main2() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := counter2{i: 0}
	done := make(chan struct{})
	for i := 0; i < 1000; i++ {
		go func() {
			c.increment()
			done <- struct{}{}
		}()
	}
	for i := 0; i < 1000; i++ {
		<-done
	}
	c.display()
}

const initialValue = -500

type counter3 struct {
	i    int64
	mu   sync.Mutex
	once sync.Once
}

func (c *counter3) increment() {
	c.once.Do(func() {
		c.i = initialValue
	})
	c.mu.Lock()
	c.i += 1
	c.mu.Unlock()
}

func (c *counter3) display() {
	fmt.Println(c.i)
}

func main3() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := counter3{i: 0}
	done := make(chan struct{})
	for i := 0; i < 1000; i++ {
		go func() {
			c.increment()
			done <- struct{}{}
		}()
	}
	for i := 0; i < 1000; i++ {
		<-done
	}
	c.display()
}

func main4() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := counter2{i: 0}
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.increment()
		}()
	}
	wg.Wait()
	c.display()
}

func (c *counter) increment2() {
	atomic.AddInt64(&c.i, 1)
}

func main5() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := counter{i: 0}
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.increment2()
		}()
	}
	wg.Wait()
	c.display()
}

func main() {
	main1()
	main2()
	main3()
	main4()
	main5()
}
