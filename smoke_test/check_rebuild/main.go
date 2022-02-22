package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("print1")
	log.Println("print2")
	log.Fatal(http.ListenAndServe(":1234", nil))
}

func demo() {
	return
}
