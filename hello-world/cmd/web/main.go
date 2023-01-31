package main

import (
	"fmt"
	"net/http"
	"github.com/example/go_webapp_demo/pkg/handlers"
)

const portNumber = ":8080"


func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}