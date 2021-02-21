package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func getFilename(link string) string {
	part := strings.Split(link, "/")
	name := part[len(part)-1]
	if len(name) == 0 {
		return fmt.Sprintf("~chaos-%d.bin", time.Now().Unix())
	}

	return name
}

func download(link string) error {
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return err
	}

	client := http.Client{}

	start := time.Now()
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	filename := getFilename(link)
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	n, err := io.CopyBuffer(f, res.Body, make([]byte, 32*1024))
	if err != nil {
		return err
	}

	fmt.Printf("Downloaded %s (%d bytes) in %.02f seconds.\n", filename, n, time.Now().Sub(start).Seconds())
	return nil
}
