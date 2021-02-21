package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		// print help
		return
	}

	if err := verifyURL(args); err != nil {
		log.Fatal(err)
	}

	const limit = 4
	wait := make(chan struct{}, limit)

	for _, arg := range args {
		wait <- struct{}{}
		go func(url string) {
			if err := download(url); err != nil {
				fmt.Println(err.Error())
			}
			<-wait
		}(arg)
	}

	for n := limit; n > 0; n-- {
		wait <- struct{}{}
	}
}
