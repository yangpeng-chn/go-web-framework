# go-web-framework

A basic web server with REST API written in Go, this application contains the following functions

1. Web server wi/wo HTTPS
2. Basic REST API with CURD operations
2. Customized Logging

## Usage

### 1. Download source code

```
git clone https://github.com/yangpeng-chn/go-web-framework.git
```

### 2. Generate self-signed certificate to enable HTTPS

1. Prepare `certs` dir

		$ cd go-web-framework
		$ mkdir certs
		
2. Genereate Private key

		$ openssl genrsa -out server.key 2048

3. Generation of self-signed(x509) public key (PEM-encodings .pem|.crt) based on the private (.key)

		$ openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650

### 3. Change configuration file to enable HTTPS

	$ vi settings/conf.json
	"EnableHTTPS": true,
	
### 4. Start web server, `go get` required packages
	
	$ go get <url>
	$ go run main.go
	
### 5. Test REST API (without HTTPS)

1. Get articles

		curl -X GET http://localhost:4201/v1/articles
	
2. Get article

		curl -X GET http://localhost:4201/v1/articles/1
		
3. Add article

		curl -X POST http://localhost:4201/v1/articles -d '{
		"id": 4,
		"title": "title4",
		"content": "content4"
		}'

4. Update articles

		curl -X PUT http://localhost:4201/v1/articles/1 -d '{"id":1,"title":"updated-title","content":"updated-content"}'
		
5. Delete article

		curl -X DELETE http://localhost:4201/v1/articles/1
		
## Other

1. Log format

		{
		 "Time": "2019-08-03 17:03:37",
		 "Message": "OK",
		 "ResponseCode": 200,
		 "Action": "GetArticlesHandler",
		 "Method": "GET",
		 "URI": "/v1/articles",
		 "RequestData": ""
		}