# Welcome to Revel

A high-productivity web framework for the [Go language](http://www.golang.org/).


### Start the web server:

   revel run myapp

### Go to http://localhost:9000/ and you'll see:

    "It works"

## Code Layout

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here
        views/        Templates directory

    messages/         Message files

    public/           Public static assets
        css/          CSS files
        js/           Javascript files
        images/       Image files

    tests/            Test suites


## Help

* The [Getting Started with Revel](http://revel.github.io/tutorial/gettingstarted.html).
* The [Revel guides](http://revel.github.io/manual/index.html).
* The [Revel sample apps](http://revel.github.io/examples/index.html).
* The [API documentation](https://godoc.org/github.com/revel/revel).

## 실행 시 오류 해결 방법
- 아래 오류는 revel new goblog 명령어 실행 후 go.mod 파일에 필요한 패키지가 누락되어 발생함
```text
no required module provides package github.com/bradfitz/gomemcache/memcache; to add it:"
```
- 일단 아래의 명령어로 memcache 모듈을 추가
```bash
$ go get github.com/bradfitz/gomemcache/memcache
# 아래 모듈은 책에서 database 설정 후 추가
$ go get github.com/mattn/go-sqlite3
```
- GoLand에서 go.mod 파일을 오른쪽 클릭 후 나오는 메뉴에서 “Go 모듈 동기화”를 실행하여 go.mod 파일을 갱신한다.

