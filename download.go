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

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(getFilename(link), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	n, err := io.CopyBuffer(f, res.Body, make([]byte, 32*1024))
	if err != nil {
		return err
	}

	fmt.Printf("Downloaded %d bytes", n)
	return nil
}
