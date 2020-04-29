# go-web-framework

A basic web backend framework with REST API written in Go, this application contains the following functions

1. Web server wi/wo HTTPS
2. Basic REST API with CURD operations
3. Customized Logging
3. Two ways of `go test`

## Usage

### 1. Download source code

```
git clone https://github.com/yangpeng-chn/go-web-framework.git
```

### 2. Change configuration file to enable HTTPS (if requried)

	$ vi settings/conf.json
	"EnableHTTPS": true,
	
### 3. Generate self-signed certificate to enable HTTPS (if changed in step 2)

1. Prepare `certs` dir

		$ cd go-web-framework
		$ mkdir certs
		
2. Genereate Private key

		$ openssl genrsa -out server.key 2048

3. Generation of self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key)

		$ openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650

### 4. Start web api

	$ export GOFLAGS=-mod=vendor
	$ go mod init github.com/yangpeng-chn/go-web-framework (go.mod generated)
	$ go mod vendor (go.mod updated, go.sum generated, vendor generated)

**4.1 dev mode**

Start with go run main.go command

	$ go run main.go
	2020-04-29 00:26:52 Listening on port 4201 ... [OK]

Or, start with development mode (hot-reload supported)

	$ docker-compose up --build
	Building go
	Step 1/4 : FROM golang:1.14
	 ---> 2421885b04da
	Step 2/4 : RUN go get github.com/oxequa/realize
	 ---> Using cache
	 ---> 2131ca7f8662
	Step 3/4 : EXPOSE 4201
	 ---> Using cache
	 ---> e5fb76b58be8
	Step 4/4 : CMD [ "realize", "start", "--run" ]
	 ---> Using cache
	 ---> c8cb5439bd09
	Successfully built c8cb5439bd09
	Successfully tagged go-web-framework_go:latest
	Starting go-web-framework_go_1 ... done
	Attaching to go-web-framework_go_1
	go_1  | [18:29:11][API] : Watching 10 file/s 7 folder/s
	go_1  | [18:29:11][API] : Build started
	go_1  | [18:29:12][API] : Build completed in 0.722 s
	go_1  | [18:29:12][API] : Running..
	go_1  | [18:29:12][API] : 2020-04-28 18:29:12 Listening on port 4201 ... [OK]

	 (stop)
	$ docker-compose down --remove-orphans --volumes

**4.2 production mode**

	$ vi docker-compose.yml
	dockerfile: Dockerfile -> dockerfile: Dockerfile.deploy

	$ docker-compose up --build

### 5. Test REST API (HTTP)

1. Get articles

		curl -X GET http://localhost:4201/v1/articles
		[{"id":1,"title":"title1","content":"content1"},{"id":2,"title":"title2","content":"content2"},{"id":3,"title":"title3","content":"content3"}]
	
2. Get article

		curl -X GET http://localhost:4201/v1/articles/1
		{"id":1,"title":"title1","content":"content1"}
		
3. Add article

		curl -X POST http://localhost:4201/v1/articles -d '{"id":4,"title": "title4","content":"content4"}'
		{"code":200,"msg":"OK"}

4. Update articles

		curl -X PUT http://localhost:4201/v1/articles/4 -d '{"id":4,"title":"updated-title","content":"updated-content"}'
		{"code":200,"msg":"OK"}
		
5. Delete article

		curl -X DELETE http://localhost:4201/v1/articles/4
		{"code":200,"msg":"OK"}

### 6. Use go test (HTTP)

	$ go test tests/article_selfserve_test.go
	ok      command-line-arguments  0.020s
	
	$ go run main.go
	$ go test tests/article_test.go
	ok      command-line-arguments  0.020s
	
## Log Format

OK

	{
	 "Time": "2019-08-03 17:03:37",
	 "Message": "OK",
	 "ResponseCode": 200,
	 "Action": "GetArticlesHandler",
	 "Method": "GET",
	 "URI": "/v1/articles",
	 "RequestData": ""
	}
		
Error

	{
	 "Time": "2019-08-14 00:08:47",
	 "Message": "article not found",
	 "ResponseCode": 400,
	 "Action": "DeleteArticleHandler",
	 "Method": "DELETE",
	 "URI": "/v1/articles/5",
	 "RequestData": ""
	}