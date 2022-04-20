package main

import (
	"hello/cmd/di"
	"net/http"
)

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(di.MyGreeterHandler))
}
