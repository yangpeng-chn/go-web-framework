# How to use `go mod`


* https://blog.framinal.life/entry/2021/04/11/013819
* https://pokuwagata.hatenablog.com/entry/2020/10/18/235429
* https://qiita.com/propella/items/e49bccc88f3cc2407745 (GO111MODULE)

## Initial status

```bash
⮀ go version
go version go1.14.2 darwin/amd64

⮀ echo $GOPATH
/Users/pengyang/repo/go-projects

⮀ ls /Users/pengyang/repo/go-projects
bin    pkg-bk src

⮀ echo $GOROOT
/usr/local/go/

⮀ pwd
/Users/pengyang/repo/go-web-framework

⮀ ls /Users/pengyang/repo/go-projects
bin    pkg-bk src

⮀ echo $GO111MODULE

⮀ ls -lh
total 28296
-rw-r--r--  1 pengyang  staff   1.7K May  2  2020 Dockerfile
-rw-r--r--  1 pengyang  staff   264B May  2  2020 Dockerfile.dev
-rw-r--r--  1 pengyang  staff   8.0K May 12  2020 README.md
drwxr-xr-x  4 pengyang  staff   128B Aug  3  2019 certs
-rw-r--r--  1 pengyang  staff   126B May 10  2020 conf.json
drwxr-xr-x  6 pengyang  staff   192B Apr 29  2020 controllers
-rw-r--r--  1 pengyang  staff   1.2K May  2  2020 docker-compose.yml
-rwxr-xr-x  1 pengyang  staff    14M May  2  2020 main
-rw-r--r--  1 pengyang  staff   832B Jul 30 23:23 main.go
drwxr-xr-x  4 pengyang  staff   128B May  2  2020 middlewares
drwxr-xr-x  6 pengyang  staff   192B May  2  2020 models
drwxr-xr-x  4 pengyang  staff   128B Aug 15  2019 tests
drwxr-xr-x  5 pengyang  staff   160B Jul 30 23:17 utils
```

## Start with `go mod init`

```bash
⮀ go mod init github.com/yangpeng-chn/go-web-framework
go: creating new go.mod: module github.com/yangpeng-chn/go-web-framework
> generate go.mod, $GOPATH/pkg/mod/cache/lock (empty)
⮀ cat go.mod
module github.com/yangpeng-chn/go-web-framework

go 1.14

⮀ ls /Users/pengyang/repo/go-projects
bin    pkg    pkg-bk src

⮀ go get github.com/joho/godotenv
go: downloading github.com/joho/godotenv v1.4.0
go: github.com/joho/godotenv upgrade => v1.4.0
> update go.mod, generate go.sum, update $GOPATH/pkg, generate $GOPATH/pkg/sumdb/

⮀ cat go.sum 
github.com/joho/godotenv v1.4.0 h1:3l4+N6zfMWnkbPEXKng2o2/MR5mSwTrBih4ZEkkz1lg=
github.com/joho/godotenv v1.4.0/go.mod h1:f4LDr5Voq0i2e/R5DDNOoa2zzDfwtkZa6DnEwAbqwq4=
 ✔ 0:00:20 [pengyang@Pengs-MacBook-Pro] ⮀ ⭠ master± ⮀ go-web-framework ⮀
⮀ cat go.mod 
module github.com/yangpeng-chn/go-web-framework

go 1.14

require github.com/joho/godotenv v1.4.0 // indirect

⮀ ls -lh /Users/pengyang/repo/go-projects/pkg/mod/github.com/joho
total 0
dr-x------  13 pengyang  staff   416B Jul 30 23:59 godotenv@v1.4.0

⮀ go mod download
> it seems it download the pkg defined in go.mod, nothing generated here as we only have godotenv in it and it was downloaded by go get

⮀ go mod vendor
> generate files in $GOPATH/pkg/mod and $GOPATH/pkg/sumdb, and copy pakcages into ./vendor dir, update go.mod and go.sum, nothing changed in $GOPATH/src

⮀ go run main.go
2022-07-31 00:10:12 Listening on port 4201 ... [OK]

⮀ rm -rf vendor

⮀ go run main.go
2022-07-31 00:10:12 Listening on port 4201 ... [OK]
> app runs without vendor dir

⮀ go clean --modcache
> $GOPATH/pkg/mod dir deleted, no change for go.mod and go.sum
```

## Deep dive `go run` and `vendor`

```bash
⮀ go run main.go
go: downloading github.com/joho/godotenv v1.4.0
go: downloading github.com/jinzhu/gorm v1.9.16
go: downloading github.com/dgrijalva/jwt-go v3.2.0+incompatible
go: downloading github.com/gorilla/mux v1.8.0
go: downloading github.com/go-sql-driver/mysql v1.5.0
go: downloading github.com/mattn/go-sqlite3 v1.14.0
go: downloading github.com/lib/pq v1.1.1
go: downloading github.com/jinzhu/inflection v1.0.0
2022-07-31 00:19:12 Listening on port 4201 ... [OK]
> $GOPATH/pkg/mod generated

⮀ go clean --modcache
> delete $GOPATH/pkg/mod/ files, sumdb dir remains

⮀ go mod download
$GOPATH/pkg/mod generated

⮀ go run -mod=mod main.go
2022-07-31 00:22:04 Listening on port 4201 ... [OK]
> use mod cache dir as source to find dependencies, it runs well even if we don't have vendor dir

⮀ go run -mod=vendor main.go
go: inconsistent vendoring in /Users/pengyang/repo/go-web-framework:
        github.com/dgrijalva/jwt-go@v3.2.0+incompatible: is explicitly required in go.mod, but not marked as explicit in vendor/modules.txt
        github.com/gorilla/mux@v1.8.0: is explicitly required in go.mod, but not marked as explicit in vendor/modules.txt
        github.com/jinzhu/gorm@v1.9.16: is explicitly required in go.mod, but not marked as explicit in vendor/modules.txt
        github.com/joho/godotenv@v1.4.0: is explicitly required in go.mod, but not marked as explicit in vendor/modules.txt
        gopkg.in/go-playground/assert.v1@v1.2.1: is explicitly required in go.mod, but not marked as explicit in vendor/modules.txt
> use vendor dir as source to find dependencies, as we don't have ./vendor dir, it throws error

⮀ go mod vendor
> copy $GOPATH/pkg to ./vendor

⮀ go run -mod=vendor main.go
2022-07-31 00:23:01 Listening on port 4201 ... [OK]

⮀ tree $GOPATH/pkg -L 3
/Users/pengyang/repo/go-projects/pkg
├── mod
│   ├── cache
│   │   └── download
│   ├── github.com
│   │   ├── !puerkito!bio
│   │   ├── andybalholm
│   │   ├── denisenkom
│   │   ├── dgrijalva
│   │   ├── erikstmartin
│   │   ├── go-sql-driver
│   │   ├── golang-sql
│   │   ├── gorilla
│   │   ├── jinzhu
│   │   ├── joho
│   │   ├── lib
│   │   └── mattn
│   ├── golang.org
│   │   └── x
│   └── gopkg.in
│       └── go-playground
└── sumdb
    └── sum.golang.org
        └── latest

22 directories, 1 file

⮀ tree $GOPATH/pkg/mod/cache -L 3
/Users/pengyang/repo/go-projects/pkg/mod/cache
└── download
    ├── github.com
    │   ├── !puerkito!bio
    │   ├── andybalholm
    │   ├── denisenkom
    │   ├── dgrijalva
    │   ├── erikstmartin
    │   ├── go-sql-driver
    │   ├── golang-sql
    │   ├── gorilla
    │   ├── jinzhu
    │   ├── joho
    │   ├── lib
    │   └── mattn
    ├── golang.org
    │   └── x
    ├── gopkg.in
    │   └── go-playground
    └── sumdb
        └── sum.golang.org

20 directories, 0 files

⮀ ls $GOPATH/pkg/mod/cache/download/github.com/joho/godotenv/@v
list           list.lock      v1.4.0.info    v1.4.0.lock    v1.4.0.mod     v1.4.0.zip     v1.4.0.ziphash

⮀ ls $GOPATH/pkg/mod/github.com/joho/godotenv@v1.4.0/
LICENCE          README.md        autoload         cmd              fixtures         go.mod           godotenv.go      godotenv_test.go renovate.json
> cache dir only has lock info while another one has source code

⮀ tree vendor -L 3
vendor
├── github.com
│   ├── dgrijalva
│   │   └── jwt-go
│   ├── go-sql-driver
│   │   └── mysql
│   ├── gorilla
│   │   └── mux
│   ├── jinzhu
│   │   ├── gorm
│   │   └── inflection
│   ├── joho
│   │   └── godotenv
│   ├── lib
│   │   └── pq
│   └── mattn
│       └── go-sqlite3
├── gopkg.in
│   └── go-playground
│       └── assert.v1
└── modules.txt

19 directories, 1 file
> not all files in pkg are copied to vendor, even tho the files in pkg folder are generated by our single project.
```

# What does `go mod tidy` do

```bash
⮀ go clean --modcache
⮀ rm go.mod go.sum
⮀ go mod init github.com/yangpeng-chn/go-web-framework

⮀ go run main.go
go: finding module for package github.com/joho/godotenv
go: finding module for package github.com/gorilla/mux
go: finding module for package github.com/jinzhu/gorm
go: finding module for package github.com/jinzhu/gorm/dialects/mysql
go: finding module for package github.com/jinzhu/gorm/dialects/sqlite
go: finding module for package github.com/dgrijalva/jwt-go
go: finding module for package github.com/jinzhu/gorm/dialects/postgres
go: downloading github.com/jinzhu/gorm v1.9.16
go: downloading github.com/joho/godotenv v1.4.0
go: downloading github.com/gorilla/mux v1.8.0
go: downloading github.com/dgrijalva/jwt-go v1.0.2
go: downloading github.com/dgrijalva/jwt-go v3.2.0+incompatible
go: found github.com/joho/godotenv in github.com/joho/godotenv v1.4.0
go: found github.com/gorilla/mux in github.com/gorilla/mux v1.8.0
go: found github.com/jinzhu/gorm in github.com/jinzhu/gorm v1.9.16
go: found github.com/dgrijalva/jwt-go in github.com/dgrijalva/jwt-go v3.2.0+incompatible
go: downloading github.com/go-sql-driver/mysql v1.5.0
go: downloading github.com/lib/pq v1.1.1
go: downloading github.com/jinzhu/inflection v1.0.0
go: downloading github.com/mattn/go-sqlite3 v1.14.0
2022-07-31 00:44:24 Listening on port 4201 ... [OK]
> go.mod updated, go.sum generated, all pkg generated in $GOPATH

⮀ go mod tidy
go: finding module for package gopkg.in/go-playground/assert.v1
go: downloading github.com/erikstmartin/go-testdb v0.0.0-20160219214506-8d10e4a1bae5
go: downloading github.com/jinzhu/now v1.0.1
go: downloading github.com/denisenkom/go-mssqldb v0.0.0-20191124224453-732737034ffd
go: downloading golang.org/x/crypto v0.0.0-20191205180655-e7c4368fe9dd
go: downloading github.com/golang-sql/civil v0.0.0-20190719163853-cb61b32ac6fe
go: downloading gopkg.in/go-playground/assert.v1 v1.2.1
go: found gopkg.in/go-playground/assert.v1 in gopkg.in/go-playground/assert.v1 v1.2.1
> update go.mod and go.sum

⮀ go clean --modcache
⮀ rm go.mod go.sum
⮀ go mod init github.com/yangpeng-chn/go-web-framework

⮀ go mod tidy
go: finding module for package github.com/jinzhu/gorm/dialects/sqlite
go: finding module for package github.com/gorilla/mux
go: finding module for package github.com/jinzhu/gorm
go: finding module for package gopkg.in/go-playground/assert.v1
go: finding module for package github.com/jinzhu/gorm/dialects/postgres
go: finding module for package github.com/jinzhu/gorm/dialects/mysql
go: finding module for package github.com/dgrijalva/jwt-go
go: finding module for package github.com/joho/godotenv
go: downloading github.com/dgrijalva/jwt-go v1.0.2
go: downloading github.com/joho/godotenv v1.4.0
go: downloading gopkg.in/go-playground/assert.v1 v1.2.1
go: downloading github.com/gorilla/mux v1.8.0
go: downloading github.com/jinzhu/gorm v1.9.16
go: downloading github.com/dgrijalva/jwt-go v3.2.0+incompatible
go: found github.com/joho/godotenv in github.com/joho/godotenv v1.4.0
go: found github.com/gorilla/mux in github.com/gorilla/mux v1.8.0
go: found github.com/jinzhu/gorm in github.com/jinzhu/gorm v1.9.16
go: found github.com/dgrijalva/jwt-go in github.com/dgrijalva/jwt-go v3.2.0+incompatible
go: found gopkg.in/go-playground/assert.v1 in gopkg.in/go-playground/assert.v1 v1.2.1
go: downloading github.com/jinzhu/inflection v1.0.0
go: downloading github.com/go-sql-driver/mysql v1.5.0
go: downloading github.com/mattn/go-sqlite3 v1.14.0
go: downloading github.com/erikstmartin/go-testdb v0.0.0-20160219214506-8d10e4a1bae5
go: downloading github.com/jinzhu/now v1.0.1
go: downloading github.com/denisenkom/go-mssqldb v0.0.0-20191124224453-732737034ffd
go: downloading github.com/lib/pq v1.1.1
go: downloading golang.org/x/crypto v0.0.0-20191205180655-e7c4368fe9dd
go: downloading github.com/golang-sql/civil v0.0.0-20190719163853-cb61b32ac6fe
> download pkg that are imported, delete pkg that are not imported, modify go.mod and go.sum

⮀ go get github.com/aws/aws-sdk-go
> update go.sum, go.mod, download pkg to $GOPATH/pkg even though we don not use it in our project.

⮀ ls $GOPATH/pkg/mod/github.com/aws
aws-sdk-go@v1.44.66

⮀ go mod tidy
⮀ ls $GOPATH/pkg/mod/github.com/aws
aws-sdk-go@v1.44.66
> update go.sum, go.mod (delete pkg info that are not imported, add pkg that are imported but not downloaded yet), actual pkg will not be deleted from $GOPATH/pkg
```

## `go build` and `go install`
```bash
⮀ go build
> executable go-web-framework generated in current working dir

⮀ ./go-web-framework
2022-07-31 00:37:36 Listening on port 4201 ... [OK]

⮀ go install
> executable go-web-framework generated in $GOPATH/bin
```

# Conclusion

1. Once we init a project, we can just issue `go run` and it will automatically download packages.
2. Run `go mod tidy` before committing your code.
3. We can create go project repositories at any locations which means it doesn't need to be $GOPATH/src