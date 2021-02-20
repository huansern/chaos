package main

import (
	"errors"
	"net/url"
)

func verifyURL(links []string) error {
	for _, link := range links {
		dl, err := url.ParseRequestURI(link)
		if err != nil {
			return err
		}

		if !(dl.Scheme == "https" || dl.Scheme == "http") {
			return errors.New("invalid url, currently only http(s) is supported")
		}
	}

	return nil
}
