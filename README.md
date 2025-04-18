# Simple CRUD operation APIs in golang gin

### Dependency

go get github.com/joho/godotenv - env setup

compiledeamon  
Watches your .go files in a directory and invokes go build if a file changed, no need to restart your server
CompileDaemon -command="./golang-gin-crud" - to execute or run go project

```
go install github.com/githubnemo/CompileDaemon@latest
```

gorm  
Used to manage db ops https://gorm.io/docs/index.html

```
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```

build docker image

```
docker build -t abhishekkadavil/golang-gin-crud:v4 .
```

push to docker hub

```
docker push abhishekkadavil/golang-gin-crud:v4
```

Execute app through docker

```
docker-compose up - run app
docker-compose down - stop app
```

Others

```
go install github.com/swaggo/swag/cmd/swag@latest
go get github.com/swaggo/swag@latest
go get github.com/swaggo/gin-swagger
go get github.com/swaggo/files
```

### Execution and swagger

```
Swag init
```

Need to do this every time when any changes happened to swagger doc
http://localhost:3000/docs/index.html
