```
brew install go
go mod init ajillo
go get github.com/99designs/gqlgen
go run github.com/99designs/gqlgen init

go get github.com/vektah/gqlparser/v2@v2.1.0
go get github.com/vektah/gqlparser/v2/ast@v2.1.0
go run ./server.go   
go clean -cache -modcache -i -r
```