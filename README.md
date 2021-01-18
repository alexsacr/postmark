# Postmark Go Client

[![GoDoc](https://godoc.org/https://github.com/alexsacr/postmark?status.svg)](https://godoc.org/github.com/alexsacr/postmark)

Under development.  You probably don't want to use this yet.

### Installation

    go get -u github.com/alexsacr/postmark

### Getting Started

```go

c := NewClient("account token", "server token")

server, rr := c.GetServer("serverID")
```
