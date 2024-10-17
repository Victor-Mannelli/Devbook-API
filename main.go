package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()

	fmt.Println("Listening to port 5000")
	r := router.NewRouter()

	log.Fatal(http.ListenAndServe(":5000", r))
}
