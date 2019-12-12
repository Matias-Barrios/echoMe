package main

import (
	"fmt"
	"github.com/Matias-Barrios/echoMe/handlers"
	"net/http"
)

const (
	port = "8888"
)

func main() {
	fmt.Println("Starting server")
	http.HandleFunc("/", handlers.PlainTextEcho)
	http.HandleFunc("/json", handlers.JsonEcho)
	http.ListenAndServe(":"+port, nil)
}
