package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("Check URL status")
		fmt.Println("ex1) urlstatus http://www.google.com")
		fmt.Println("ex2) urlstatus http://www.google.com http://www.naver.com")
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
	resp, err := http.Get(url)
	rcode := "ERR"
	if err != nil {
		ch <- fmt.Sprintf(" %3s %3d %s", rcode, 0, url)
		return
	}
	snum := resp.StatusCode
	switch snum {
		case 200 : rcode="OK" //요청성공 
		case 201 : rcode="OK" //요청성공 및 새 리소스 작성
		case 202 : rcode="OK" //요청을 접수, 아직 처리하지 않았음.
		case 203 : rcode="OK" //요청접수. 다른소스에서 수신된 정보를 제공.
		case 204 : rcode="OK" //요청접수. 콘텐츠 없음.
		case 205 : rcode="OK" //요청성공. 콘첸츠 표시X.
		case 206 : rcode="OK" //서버가 일부요청만 성공처리.
		case 300 : rcode="OK" //3XX리다이렉션 이슈.
		case 301 : rcode="OK" //https://ko.wikipedia.org/wiki/HTTP_상태_코드
		case 302 : rcode="OK"
		case 303 : rcode="OK"
		case 304 : rcode="OK"
		case 305 : rcode="OK"
		case 306 : rcode="OK"
		case 401 : rcode="OK" //권한이 없음. 인증필요.
	}
	ch <- fmt.Sprintf(" %3s %3d %s", rcode, snum, url)
}
