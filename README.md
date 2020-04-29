# go-web-framework

A basic web backend framework with REST API written in Go, this application contains the following functions

1. Web server wi/wo HTTPS
2. Basic REST API with CURD operations
3. Manipulate data in memroy (articles) and mysql database (posts)
4. Authentication for `posts` resource
5. Dev mode supporting hot-reload with `realize`, production mode by deploying compiled binary to container
6. Run webapi, database and phpmyadmin in different containers
7. Customized logging
8. Two ways of `go test` (more tests to be added)

## Usage

### 1. Download source code

```
git clone https://github.com/yangpeng-chn/go-web-framework.git
```

### 2. Change configuration file to enable HTTPS (if requried)

	$ vi conf/conf.json
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

	 (stop)
	$ docker-compose down --remove-orphans --volumes

### 5. Test REST API (HTTP)

Articles are stored in memory while posts are stored in database

1. Get articles

		curl -X GET http://localhost:4201/v1/articles
		[{"id":1,"title":"title1","content":"content1"},{"id":2,"title":"title2","content":"content2"},{"id":3,"title":"title3","content":"content3"}]
	
2. Get article

		curl -X GET http://localhost:4201/v1/articles/1
		{"id":1,"title":"title1","content":"content1"}
		
3. Add article

		curl -X POST http://localhost:4201/v1/articles -d '{"id":4,"title": "title4","content":"content4"}'
		{"id":4,"title":"title4","content":"content4"}

4. Update articles

		curl -X PUT http://localhost:4201/v1/articles/4 -d '{"id":4,"title":"updated-title","content":"updated-content"}'
		{"id":4,"title":"updated-title","content":"updated-content"}
		
5. Delete article

		curl -X DELETE http://localhost:4201/v1/articles/4
		{"code":200,"msg":"OK"}

		curl -X DELETE http://localhost:4201/v1/articles/5              
		{"error":"article not found"}

6. Get posts

		curl -X GET http://localhost:4201/v1/posts 
		[{"id":1,"title":"Title 1","content":"Hello world 1","author":{"id":1,"nickname":"Yang","email":"yang@gmail.com","password":"password","created_at":"2020-04-29T14:54:36Z","updated_at":"2020-04-29T14:54:36Z"},"author_id":1,"created_at":"2020-04-29T14:54:36Z","updated_at":"2020-04-29T14:54:36Z"},{"id":2,"title":"Title 2","content":"Hello world 2","author":{"id":2,"nickname":"Martin Luther","email":"luther@gmail.com","password":"password","created_at":"2020-04-29T14:54:36Z","updated_at":"2020-04-29T14:54:36Z"},"author_id":2,"created_at":"2020-04-29T14:54:36Z","updated_at":"2020-04-29T14:54:36Z"}]

7. Get post

		curl -X GET http://localhost:4201/v1/posts/1
		{"id":1,"title":"Title 1","content":"Hello world 1","author":{"id":1,"nickname":"Yang","email":"yang@gmail.com","password":"password","created_at":"2020-04-29T14:54:36Z","updated_at":"2020-04-29T14:54:36Z"},"author_id":1,"created_at":"2020-04-29T14:54:36Z","updated_at":"2020-04-29T14:54:36Z"}

8. Add post

		curl -X POST http://localhost:4201/v1/posts -d '{"id":3,"title":"title 3","content":"content 3"} ...'

9.  Update post

		curl -X PUT http://localhost:4201/v1/posts/1 -d '{"id":1,"title":"updated-title","content":"updated-content"}'
		{"error":"Unauthorized"}

10.  Delete post

		curl -X DELETE http://localhost:4201/v1/posts/1
		{"error":"Unauthorized"}

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