## Go Withdraw API

Withdraw REST API service using Go and Gin.

## Running

The version used on development is `go1.15.6`

```bash
$ git clone git@github.com:raphaelkross/go-withdraw-api.git

$ cd go-withdraw-api

$ go run main.go
```

## [GET] v1/withdraw/:amount

Get the minimum amount of bills given a int amount. Bills available: 50, 10, 5 and 1.

```
curl -X GET \
  http://localhost:8080/v1/withdraw/187 \
  -H 'Host: localhost:8080'
```

## Unit Test

```bash
# run all unit tests
$ go test ./...
```
