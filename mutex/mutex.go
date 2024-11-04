package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
	// Introducing mutex to avoid race condition
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		p.mu.Unlock()
	}()

	p.mu.Lock()
	p.views += 1
}

func (p *post) dec(wg *sync.WaitGroup) {
	defer wg.Done()
	p.views -= 1
}

func main() {

	var wg sync.WaitGroup

	post1 := post{views: 0}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go post1.inc(&wg)
	}

	wg.Wait()
	fmt.Println(post1.views)
}
