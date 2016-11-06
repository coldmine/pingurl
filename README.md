# PINGURL
- 해당 URL 웹서버가 살아있는지 체크하기 위한 유틸리티.
- 모든 인수에 대한 URL을 동시에 처리합니다. = 속도가 빠릅니다.

#### Install
```
go get -u github.com/coldmine/pingurl
```

#### How to use
```
$ pingurl http://www.google.com
$ pingurl http://www.naver.com http://www.google.com //여러개의 URL을 설정할 수 있습니다.
```

#### 제작의도
- `$ ping www.naver.com`처럼 ping을 막은 사이트의 경우 살아있는지 체크하기 힘들다.

#### License : GPLv3

