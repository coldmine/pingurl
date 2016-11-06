# PINGURL
- 해당 URL 서버가 살아있는지 체크하기 위해서 제작된 코드입니다.
- 위 이슈로 401 상태코드를 OK로 처리합니다.
- 모든 인수에 대한 URL을 동시에 처리합니다.

#### Install
```
go get -u github.com/coldmine/pingurl
```

#### How to use
```
$ pingurl http://www.google.com
$ pingurl http://www.naver.com http://www.google.com //여러개의 URL을 설정할 수 있습니다.
$ pingurl 8.8.8.8 //구글DNS서버는 http 상태코드를 반환하지 않기 때문에 조금 시간이 걸린후 ERR로 실패됩니다.
```

#### License : GPLv3

