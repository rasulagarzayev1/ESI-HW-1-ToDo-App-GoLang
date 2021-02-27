package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HashimovH/todo/backend/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server")
	log.Fatal(http.ListenAndServe(":8080", r))
}
