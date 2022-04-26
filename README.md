# simple_server_golang
This is a simple server implementation with golang.

## ⚙️ Getting Started
```
$ git clone https://github.com/filipebafica/simple_server_golang.git
$ cd simple_server_golang
$ go mod tidy
```
## 🎈 How to Use
Run the server on a terminal:
```
$ go run src/main.go
```
On a different terminal send a request:
```
$ curl "localhost:9090"
>>> Hello

$ curl "localhost:9090/goodbye"
>>> Byeee
```
