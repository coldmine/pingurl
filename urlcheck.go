package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

var Version string

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Printf("URL Check %s\n", Version)
		fmt.Println("ex1) $ urlcheck http://www.google.com")
		fmt.Println("ex2) $ urlcheck http://www.google.com http://www.naver.com")
		os.Exit(1)
	}
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
}

func fetch(url string, ch chan<- string) {
	if !(strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")) {
		ch <- fmt.Sprintf(` ERR     %s : Begin with "http://" or "https://".`, url)
		return
	}
	resp, err := http.Get(url)
	rcode := "OK"
	if err != nil {
		ch <- fmt.Sprintf(" ERR %3d %s : %s", 0, url, err)
		return
	}
	snum := resp.StatusCode
	if snum > 499 {
		rcode = "ERR"
	}
	ch <- fmt.Sprintf(" %3s %3d %s", rcode, snum, url)
}
