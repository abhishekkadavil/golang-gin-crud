go get github.com/joho/godotenv - env setup

compiledeamon
go install github.com/githubnemo/CompileDaemon@latest - Watches your .go files in a directory and invokes go build if a file changed, no need to restart your server
CompileDaemon -command="./golang-gin-crud" - to execute or run go project

gorm - used to manage db ops
https://gorm.io/docs/index.html
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres