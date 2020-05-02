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

```bash
$ vi conf/conf.json
"EnableHTTPS": true,
```
	
### 3. Generate self-signed certificate to enable HTTPS (if changed in step 2)

1. Prepare `certs` dir

	```bash
	$ cd go-web-framework
	$ mkdir certs
	$ cd certs
	```
		
2. Genereate Private key

	```bash
	$ openssl genrsa -out server.key 2048
	```

3. Generation of self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key)

	```bash
	$ openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
	```

### 4. Start web api

```bash
$ export GOFLAGS=-mod=vendor
$ export GO111MODULE=on 
$ go mod init github.com/yangpeng-chn/go-web-framework (go.mod generated)
$ go mod vendor (go.mod updated, go.sum generated, vendor generated)
```

**Dev mode**

1. Start service with `go run main.go` (without hot-reload and docker)

	```bash
	$ go run main.go
	2020-04-29 00:26:52 Listening on port 4201 ... [OK]
	```

2. Start service with `realize` with hot-reload and without docker, database not available

	```bash
	$ vi conf/conf.json
	"UseDatabase": false,

	$ GO111MODULE=off go get github.com/oxequa/realize
	$ vi ~/.zprofile
	$ source ~/.zprofile
	---
	export PATH=$PATH:$GOPATH/bin
	---
	$ which realize

	$ realize start --run # also works without --run
	[20:45:52][API] : Watching 157 file/s 45 folder/s
	[20:45:52][API] : Build started
	[20:45:52][API] : Build completed in 0.546 s
	[20:45:52][API] : Running..
	[20:45:53][API] : 2020-05-01 20:45:53 Listening on port 4201 ... [OK]
	```

3. Start with docker-compose (hot-reload supported)

	```bash
	$ docker-compose up --build
	Building app
	Step 1/7 : FROM golang:1.14
	---> 2421885b04da
	Step 2/7 : RUN go get github.com/oxequa/realize
	---> Using cache
	---> 2131ca7f8662
	Step 3/7 : ENV APP_HOME /app
	---> Using cache
	---> a6a9a670c9cb
	Step 4/7 : RUN mkdir -p $APP_HOME
	---> Using cache
	---> c2467a23fca1
	Step 5/7 : WORKDIR $APP_HOME
	---> Using cache
	---> 0356ac1555ac
	Step 6/7 : EXPOSE 4201
	---> Using cache
	---> 7879635444d2
	Step 7/7 : CMD [ "realize", "start", "--run" ]
	---> Using cache
	---> cc1b8be1f056

	Successfully built cc1b8be1f056
	Successfully tagged go-web-framework_app:latest
	Starting db_mysql ... done
	Starting phpmyadmin ... done
	Starting full_app   ... done
	Attaching to db_mysql, phpmyadmin, full_app
	...
	full_app           | [11:46:28][API] : Watching 157 file/s 45 folder/s
	full_app           | [11:46:28][API] : Build started
	db_mysql           | 2020-05-01T11:46:28.154363Z 0 [Note] Event Scheduler: Loaded 0 events
	db_mysql           | 2020-05-01T11:46:28.156387Z 0 [Note] mysqld: ready for connections.
	db_mysql           | Version: '5.7.29'  socket: '/var/run/mysqld/mysqld.sock'  port: 3306  MySQL Community Server (GPL)
	full_app           | [11:46:40][API] : Build completed in 12.178 s
	full_app           | [11:46:40][API] : Running..
	full_app           | [11:46:40][API] : 2020-05-01 11:46:40 Listening on port 4201 ... [OK]

	 (stop)
	$ docker-compose down --remove-orphans --volumes
	```

**Production mode**

```bash
$ vi docker-compose.yml
dockerfile: Dockerfile.dev -> dockerfile: Dockerfile

$ docker-compose up --build

 (stop)
$ docker-compose down --remove-orphans --volumes
```

### 5. Test REST API (HTTP)

Articles are stored in memory while posts are stored in database

1. Get articles
   
	```bash
	curl -X GET http://localhost:4201/v1/articles
	[{"id":1,"title":"title1","content":"content1"},{"id":2,"title":"title2","content":"content2"},{"id":3,"title":"title3","content":"content3"}]
	```
	
2. Get article

	```bash
	curl -X GET http://localhost:4201/v1/articles/1
	{"id":1,"title":"title1","content":"content1"}
	```
		
3. Add article

	```bash
	curl -X POST http://localhost:4201/v1/articles -d '{"id":4,"title": "title4","content":"content4"}'{"id":4,"title":"title4","content":"content4"}
	```

4. Update articles

	```bash
	curl -X PUT http://localhost:4201/v1/articles/4 -d '{"id":4,"title":"updated-title","content":"updated-content"}'{"id":4,"title":"updated-title","content":"updated-content"}
	```
		
5. Delete article

	```bash
	curl -X DELETE http://localhost:4201/v1/articles/4
	{"code":200,"msg":"OK"}

	curl -X DELETE http://localhost:4201/v1/articles/5              
	{"error":"article not found"}
	```

6. Get posts

	```bash
	curl -X GET http://localhost:4201/v1/posts 
	[{"id":1,"title":"Title 1","content":"Hello world 1","author":{"id":1,"nickname":"Yang","email":"yang@gmail.com","password":"password","created_at":"2020-04-29T14:54:36Z","updated_at":"2020-04-29T14:54:36Z"},"author_id":1,"created_at":"2020-04-29T14:54:36Z","updated_at":"2020-04-29T14:54:36Z"},{"id":2,"title":"Title 2","content":"Hello world 2","author":{"id":2,"nickname":"Martin Luther","email":"luther@gmail.com","password":"password","created_at":"2020-04-29T14:54:36Z","updated_at":"2020-04-29T14:54:36Z"},"author_id":2,"created_at":"2020-04-29T14:54:36Z","updated_at":"2020-04-29T14:54:36Z"}]
	```

7. Get post

	```bash
	curl -X GET http://localhost:4201/v1/posts/1
	{"id":1,"title":"Title 1","content":"Hello world 1","author":{"id":1,"nickname":"Yang","email":"yang@gmail.com","password":"password","created_at":"2020-04-29T14:54:36Z","updated_at":"2020-04-29T14:54:36Z"},"author_id":1,"created_at":"2020-04-29T14:54:36Z","updated_at":"2020-04-29T14:54:36Z"}
	```

8. Add post

	```bash
	curl -X POST http://localhost:4201/v1/posts -d '{"id":3,"title":"title 3","content":"content 3"} ...'
	```

9.  Update post

	```bash
	curl -X PUT http://localhost:4201/v1/posts/1 -d '{"id":1,"title":"updated-title","content":"updated-content"}'{"error":"Unauthorized"}
	```

9.  Delete post

	```bash
	curl -X DELETE http://localhost:4201/v1/posts/1
	{"error":"Unauthorized"}
	```

### 6. Use go test (HTTP)

```bash
$ go test tests/article_selfserve_test.go
ok      command-line-arguments  0.020s
	
$ go run main.go
$ go test tests/article_test.go
ok      command-line-arguments  0.020s
```
	
## Log Format

OK

```bash
{
 "Time": "2019-08-03 17:03:37",
 "Message": "OK",
 "ResponseCode": 200,
 "Action": "GetArticlesHandler",
 "Method": "GET",
 "URI": "/v1/articles",
 "RequestData": ""
}
```

Error

```bash
{
 "Time": "2019-08-14 00:08:47",
 "Message": "article not found",
 "ResponseCode": 400,
 "Action": "DeleteArticleHandler",
 "Method": "DELETE",
 "URI": "/v1/articles/5",
 "RequestData": ""
}
```

## Note

1. [go mod](https://medium.com/@petomalina/using-go-mod-download-to-speed-up-golang-docker-builds-707591336888)

2. Contents in container when using `Dockerfile.dev` and `docker-compose.yml`

```bash
⮀ docker ps -a
CONTAINER ID        IMAGE                   COMMAND                  CREATED             STATUS              PORTS                               NAMES
32e5d0bb8d85        go-web-framework_app    "realize start --run"    50 seconds ago      Up 49 seconds       0.0.0.0:4201->4201/tcp              full_app
4360effe1a7d        phpmyadmin/phpmyadmin   "/docker-entrypoint.…"   4 minutes ago       Up 49 seconds       0.0.0.0:9090->80/tcp                phpmyadmin
555d1f9a5228        mysql:5.7               "docker-entrypoint.s…"   4 minutes ago       Up 50 seconds       0.0.0.0:3306->3306/tcp, 33060/tcp   db_mysql
⮀ docker exec -it 32e5d0bb8d85 sh
# pwd; ls
/build
Dockerfile  Dockerfile.dev  README.md  certs  conf.json  controllers  docker-compose.yml  go.mod  go.sum  main  main.go  middlewares  models  tests  utils
```

3. Other command

```bash
# docker build -t myapp-deploy -fDockerfile .
# docker run -it -p 4201:4201 myapp-deploy
```
