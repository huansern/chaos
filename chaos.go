package main

import (
	"context"
	"fmt"
	"sync"
)

const MaxConcurrentDownloads = 4

func chaos(ctx context.Context, urls []string) {
	wg := &sync.WaitGroup{}
	task := make(chan string)

	n := len(urls)
	if n > MaxConcurrentDownloads {
		n = MaxConcurrentDownloads
	}
	for ; n > 0; n-- {
		wg.Add(1)
		go worker(ctx, wg, task)
	}

	go func() {
		defer close(task)
		for _, url := range urls {
			select {
			case <-ctx.Done():
				break
			case task <- url:
			}
		}
	}()

	wg.Wait()
}

func worker(ctx context.Context, wg *sync.WaitGroup, urls <-chan string) {
	defer wg.Done()
	for url := range urls {
		if err := download(ctx, url); err != nil {
			fmt.Println(err.Error())
		}
	}
}
