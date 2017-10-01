package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// GetURL : a function used to get the url
func GetURL() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "https://") {
			url = "https://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("http-get: %v, Status Code: %d\n", resp.Status, resp.StatusCode)

		//b, err := ioutil.ReadAll(resp.Body)
		_, err = io.Copy(os.Stdout, resp.Body)
		// copy copies from src to src.

		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		defer resp.Body.Close()
		// fmt.Printf("%s", b)
	}
}
