package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
)

const MaxConcurrentDownloads = 4

func chaos(ctx context.Context, urls []string) {
	wg := &sync.WaitGroup{}
	task := make(chan string)
	client := &http.Client{}

	n := len(urls)
	if n > MaxConcurrentDownloads {
		n = MaxConcurrentDownloads
	}
	for ; n > 0; n-- {
		wg.Add(1)
		go worker(ctx, wg, client, task)
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

func worker(ctx context.Context, wg *sync.WaitGroup, client *http.Client, urls <-chan string) {
	defer wg.Done()
	for url := range urls {
		if err := download(ctx, client, url); err != nil {
			fmt.Println(err.Error())
		}
	}
}
