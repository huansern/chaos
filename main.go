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

	for _, link := range args {
		if err := download(link); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}
}
