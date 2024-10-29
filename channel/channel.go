package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	chanslice := []<-chan interface{}{
		sig(2 * time.Hour),
		sig(5 * time.Minute),
		sig(2 * time.Second),
		sig(1 * time.Hour),
		sig(1 * time.Minute),
	}

	endChan := or(chanslice...)
	<-endChan

	fmt.Printf("fone after %v", time.Since(start))
}

func sig(after time.Duration) <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		time.Sleep(after)
	}()
	return ch
}

func or(channels ...<-chan interface{}) <-chan interface{} {

	if len(channels) == 0 {
		ch := make(chan interface{})
		close(ch)
		return ch
	}

	wg := sync.WaitGroup{}
	on := sync.Once{}
	orchan := make(chan interface{})

	go func() {
		for _, c := range channels {
			wg.Add(1)
			go func(c <-chan interface{}) {
				defer wg.Done()

				for it := range c {
					orchan <- it
				}
				on.Do(func() {
					close(orchan)
				})
			}(c)
		}
		wg.Wait()
	}()

	return orchan
}
